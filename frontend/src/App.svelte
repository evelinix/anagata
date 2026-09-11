<script lang="ts">
  import logo from './assets/images/logo-universal.png';
  import { Greet, GetVersionInfo } from '../wailsjs/go/app/App';
  import { getTheme, toggleTheme, initTheme } from './lib/stores/theme';
  import { t, getLocale, setLocale } from './lib/i18n/index';
  import { Button, Card, IconButton, Typography } from './lib/components';
  import {
    IconWorld,
    IconSun,
    IconMoon,
    IconLayoutDashboard,
    IconSearch,
    IconSettings,
    IconInfoCircle,
    IconChartBar,
    IconAlertTriangle,
    IconFileText,
    IconBell,
  } from './lib/components/icons';

  let resultText = $state('');
  let name = $state('');
  let locale = $state(getLocale());
  let localeVersion = $state(0);
  let currentTheme = $state(getTheme());
  let versionInfo = $state<Record<string, string>>({});

  $effect(() => {
    initTheme();
    currentTheme = getTheme();
    GetVersionInfo().then((info: Record<string, string>) => {
      versionInfo = info;
    });
    const onLocale = (e: Event) => {
      locale = (e as CustomEvent).detail;
      localeVersion++;
    };
    const onTheme = (e: Event) => {
      currentTheme = (e as CustomEvent).detail;
    };
    window.addEventListener('localechange', onLocale);
    window.addEventListener('themechange', onTheme);
    return () => {
      window.removeEventListener('localechange', onLocale);
      window.removeEventListener('themechange', onTheme);
    };
  });

  function greet() {
    Greet(name).then((result: string) => {
      resultText = result;
    });
  }

  function toggleLocale() {
    const next = locale === 'en' ? 'id' : 'en';
    setLocale(next);
  }

  function handleToggleTheme() {
    toggleTheme();
  }

  function tl(key: string): string {
    void localeVersion;
    return t(key);
  }
</script>

<div
  class="min-h-screen flex flex-col"
  style="background-color: var(--bg-mica); color: var(--text-primary);"
>
  <!-- Header -->
  <header
    class="flex items-center justify-between px-5 py-3"
    style="border-bottom: 1px solid var(--stroke-subtle); background-color: var(--bg-layer-alt);"
  >
    <div class="flex items-center gap-3">
      <img src={logo} class="w-6 h-6" alt="" />
      <Typography variant="body-strong">{tl('app.title')}</Typography>
    </div>
    <div class="flex items-center gap-1">
      <IconButton variant="subtle" size="sm" label="Toggle language" onclick={toggleLocale}>
        <IconWorld size={16} stroke={2} />
      </IconButton>
      <IconButton variant="subtle" size="sm" label={tl('theme.toggle')} onclick={handleToggleTheme}>
        {#if currentTheme === 'dark'}
          <IconSun size={16} stroke={2} />
        {:else}
          <IconMoon size={16} stroke={2} />
        {/if}
      </IconButton>
    </div>
  </header>

  <!-- Main Content -->
  <main class="flex-1 flex flex-col items-center justify-center px-6 py-10">
    <!-- Hero -->
    <div class="flex flex-col items-center mb-8">
      <img src={logo} class="w-16 h-16 mb-4" alt="" style="filter: drop-shadow(var(--shadow-4));" />
      <Typography variant="title" as="h1" weight="bold" class="mb-1">
        {tl('app.title')}
      </Typography>
      <Typography
        variant="body"
        weight="semibold"
        color="accent"
        align="center"
        uppercase
        class="mb-2"
        style="letter-spacing: 0.08em;"
      >
        {tl('app.subtitle')}
      </Typography>
      <Typography variant="body" color="secondary" align="center" style="max-width: 28rem;">
        {tl('app.description')}
      </Typography>
    </div>

    <!-- Status Cards -->
    <div class="grid grid-cols-2 md:grid-cols-4 gap-3 mb-8 w-full" style="max-width: 42rem;">
      <Card variant="default" padding>
        <div class="flex flex-col items-center gap-2">
          <div
            style="width: 8px; height: 8px; border-radius: var(--radius-full); background-color: var(--status-success);"
          ></div>
          <Typography variant="caption" color="tertiary" uppercase>{tl('status.system')}</Typography
          >
          <Typography variant="body-strong" color="success">{tl('status.online')}</Typography>
        </div>
      </Card>
      <Card variant="default" padding>
        <div class="flex flex-col items-center gap-2">
          <div
            style="width: 8px; height: 8px; border-radius: var(--radius-full); background-color: var(--status-success);"
          ></div>
          <Typography variant="caption" color="tertiary" uppercase
            >{tl('status.threats')}</Typography
          >
          <Typography variant="body-strong" color="success">{tl('status.noThreats')}</Typography>
        </div>
      </Card>
      <Card variant="default" padding>
        <div class="flex flex-col items-center gap-2">
          <div
            style="width: 8px; height: 8px; border-radius: var(--radius-full); background-color: var(--status-info);"
          ></div>
          <Typography variant="caption" color="tertiary" uppercase
            >{tl('status.monitors')}</Typography
          >
          <Typography variant="body-strong" color="info">12</Typography>
        </div>
      </Card>
      <Card variant="default" padding>
        <div class="flex flex-col items-center gap-2">
          <div
            style="width: 8px; height: 8px; border-radius: var(--radius-full); background-color: var(--status-neutral);"
          ></div>
          <Typography variant="caption" color="tertiary" uppercase
            >{tl('status.lastScan')}</Typography
          >
          <Typography variant="body-strong" color="secondary">{tl('status.never')}</Typography>
        </div>
      </Card>
    </div>

    <!-- Quick Actions -->
    <div class="flex gap-3 mb-10">
      <Button variant="primary" size="md" onclick={greet}>
        <IconLayoutDashboard size={16} stroke={2} />
        {tl('actions.dashboard')}
      </Button>
      <Button variant="secondary" size="md" onclick={greet}>
        <IconSearch size={16} stroke={2} />
        {tl('actions.scan')}
      </Button>
      <Button variant="outline" size="md" onclick={greet}>
        <IconSettings size={16} stroke={2} />
        {tl('actions.settings')}
      </Button>
    </div>

    <!-- Greet Result -->
    {#if resultText}
      <div
        class="mb-8 flex items-center gap-3"
        style="padding: 12px 16px; border-radius: var(--radius-md); background-color: var(--status-info-bg); border: 1px solid var(--status-info); box-shadow: var(--shadow-2);"
      >
        <IconInfoCircle size={16} stroke={2} style="flex-shrink: 0; color: var(--status-info);" />
        <Typography variant="body">{resultText}</Typography>
      </div>
    {/if}

    <!-- Features Grid -->
    <div class="w-full" style="max-width: 42rem;">
      <Typography
        variant="caption-strong"
        color="tertiary"
        uppercase
        align="center"
        class="mb-4"
        style="letter-spacing: 0.08em;"
      >
        {tl('features.title')}
      </Typography>
      <div class="grid grid-cols-2 md:grid-cols-4 gap-3">
        <Card variant="default" padding>
          <div class="flex flex-col items-center gap-2">
            <IconChartBar size={20} stroke={2} style="color: var(--accent-default);" />
            <Typography variant="caption-strong" align="center"
              >{tl('features.monitoring')}</Typography
            >
          </div>
        </Card>
        <Card variant="default" padding>
          <div class="flex flex-col items-center gap-2">
            <IconAlertTriangle size={20} stroke={2} style="color: var(--status-critical);" />
            <Typography variant="caption-strong" align="center">{tl('features.threats')}</Typography
            >
          </div>
        </Card>
        <Card variant="default" padding>
          <div class="flex flex-col items-center gap-2">
            <IconFileText size={20} stroke={2} style="color: var(--accent-default);" />
            <Typography variant="caption-strong" align="center"
              >{tl('features.analytics')}</Typography
            >
          </div>
        </Card>
        <Card variant="default" padding>
          <div class="flex flex-col items-center gap-2">
            <IconBell size={20} stroke={2} style="color: var(--status-attention);" />
            <Typography variant="caption-strong" align="center">{tl('features.alerts')}</Typography>
          </div>
        </Card>
      </div>
    </div>
  </main>

  <!-- Footer -->
  <footer
    class="flex items-center justify-between px-5 py-2"
    style="border-top: 1px solid var(--stroke-subtle); background-color: var(--bg-layer-alt);"
  >
    <Typography variant="caption" color="tertiary">
      v{versionInfo.version || '0.1.0'}
    </Typography>
    <Typography variant="caption" color="tertiary">
      &copy; {new Date().getFullYear()}
      {tl('footer.copyright')}
    </Typography>
  </footer>
</div>
