<script lang="ts">
  import logo from './assets/images/logo-universal.png';
  import { Greet, GetVersionInfo } from '../wailsjs/go/app/App';
  import { getTheme, toggleTheme, initTheme } from './lib/stores/theme';
  import { t, getLocale, setLocale } from './lib/i18n/index';
  import { Button, Card, IconButton } from './lib/components';
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
  } from '@tabler/icons-svelte';

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
      <span
        class="font-semibold"
        style="font-size: var(--text-body); line-height: var(--text-body-lh); color: var(--text-primary);"
        >{tl('app.title')}</span
      >
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
      <h1
        style="font-size: var(--text-title); line-height: var(--text-title-lh); font-weight: var(--weight-bold); color: var(--text-primary);"
        class="mb-1"
      >
        {tl('app.title')}
      </h1>
      <p
        style="font-size: var(--text-body); line-height: var(--text-body-lh); font-weight: var(--weight-semibold); color: var(--accent-default); letter-spacing: 0.08em; text-transform: uppercase;"
        class="mb-2"
      >
        {tl('app.subtitle')}
      </p>
      <p
        style="font-size: var(--text-body); line-height: var(--text-body-lh); color: var(--text-secondary); max-width: 28rem; text-align: center;"
      >
        {tl('app.description')}
      </p>
    </div>

    <!-- Status Cards -->
    <div class="grid grid-cols-2 md:grid-cols-4 gap-3 mb-8 w-full" style="max-width: 42rem;">
      <Card variant="default" padding>
        <div class="flex flex-col items-center gap-2">
          <div
            style="width: 8px; height: 8px; border-radius: var(--radius-full); background-color: var(--status-success);"
          ></div>
          <span
            style="font-size: var(--text-caption); line-height: var(--text-caption-lh); color: var(--text-tertiary); text-transform: uppercase; letter-spacing: 0.06em;"
            >{tl('status.system')}</span
          >
          <span
            style="font-size: var(--text-body); line-height: var(--text-body-lh); font-weight: var(--weight-semibold); color: var(--status-success-text);"
            >{tl('status.online')}</span
          >
        </div>
      </Card>
      <Card variant="default" padding>
        <div class="flex flex-col items-center gap-2">
          <div
            style="width: 8px; height: 8px; border-radius: var(--radius-full); background-color: var(--status-success);"
          ></div>
          <span
            style="font-size: var(--text-caption); line-height: var(--text-caption-lh); color: var(--text-tertiary); text-transform: uppercase; letter-spacing: 0.06em;"
            >{tl('status.threats')}</span
          >
          <span
            style="font-size: var(--text-body); line-height: var(--text-body-lh); font-weight: var(--weight-semibold); color: var(--status-success-text);"
            >{tl('status.noThreats')}</span
          >
        </div>
      </Card>
      <Card variant="default" padding>
        <div class="flex flex-col items-center gap-2">
          <div
            style="width: 8px; height: 8px; border-radius: var(--radius-full); background-color: var(--status-info);"
          ></div>
          <span
            style="font-size: var(--text-caption); line-height: var(--text-caption-lh); color: var(--text-tertiary); text-transform: uppercase; letter-spacing: 0.06em;"
            >{tl('status.monitors')}</span
          >
          <span
            style="font-size: var(--text-body); line-height: var(--text-body-lh); font-weight: var(--weight-semibold); color: var(--status-info-text);"
            >12</span
          >
        </div>
      </Card>
      <Card variant="default" padding>
        <div class="flex flex-col items-center gap-2">
          <div
            style="width: 8px; height: 8px; border-radius: var(--radius-full); background-color: var(--status-neutral);"
          ></div>
          <span
            style="font-size: var(--text-caption); line-height: var(--text-caption-lh); color: var(--text-tertiary); text-transform: uppercase; letter-spacing: 0.06em;"
            >{tl('status.lastScan')}</span
          >
          <span
            style="font-size: var(--text-body); line-height: var(--text-body-lh); font-weight: var(--weight-semibold); color: var(--text-secondary);"
            >{tl('status.never')}</span
          >
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
        style="padding: 12px 16px; border-radius: var(--radius-md); background-color: var(--status-info-bg); border: 1px solid var(--status-info); font-size: var(--text-body); line-height: var(--text-body-lh); color: var(--status-info-text); box-shadow: var(--shadow-2);"
      >
        <IconInfoCircle size={16} stroke={2} style="flex-shrink: 0; color: var(--status-info);" />
        {resultText}
      </div>
    {/if}

    <!-- Features Grid -->
    <div class="w-full" style="max-width: 42rem;">
      <h2
        style="font-size: var(--text-caption); line-height: var(--text-caption-lh); font-weight: var(--weight-semibold); color: var(--text-tertiary); text-transform: uppercase; letter-spacing: 0.08em;"
        class="text-center mb-4"
      >
        {tl('features.title')}
      </h2>
      <div class="grid grid-cols-2 md:grid-cols-4 gap-3">
        <Card variant="default" padding>
          <div class="flex flex-col items-center gap-2">
            <IconChartBar size={20} stroke={2} style="color: var(--accent-default);" />
            <span
              style="font-size: var(--text-caption); line-height: var(--text-caption-lh); font-weight: var(--weight-semibold); color: var(--text-primary); text-align: center;"
              >{tl('features.monitoring')}</span
            >
          </div>
        </Card>
        <Card variant="default" padding>
          <div class="flex flex-col items-center gap-2">
            <IconAlertTriangle size={20} stroke={2} style="color: var(--status-critical);" />
            <span
              style="font-size: var(--text-caption); line-height: var(--text-caption-lh); font-weight: var(--weight-semibold); color: var(--text-primary); text-align: center;"
              >{tl('features.threats')}</span
            >
          </div>
        </Card>
        <Card variant="default" padding>
          <div class="flex flex-col items-center gap-2">
            <IconFileText size={20} stroke={2} style="color: var(--accent-default);" />
            <span
              style="font-size: var(--text-caption); line-height: var(--text-caption-lh); font-weight: var(--weight-semibold); color: var(--text-primary); text-align: center;"
              >{tl('features.analytics')}</span
            >
          </div>
        </Card>
        <Card variant="default" padding>
          <div class="flex flex-col items-center gap-2">
            <IconBell size={20} stroke={2} style="color: var(--status-attention);" />
            <span
              style="font-size: var(--text-caption); line-height: var(--text-caption-lh); font-weight: var(--weight-semibold); color: var(--text-primary); text-align: center;"
              >{tl('features.alerts')}</span
            >
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
    <span
      style="font-size: var(--text-caption); line-height: var(--text-caption-lh); color: var(--text-tertiary);"
    >
      v{versionInfo.version || '0.1.0'}
    </span>
    <span
      style="font-size: var(--text-caption); line-height: var(--text-caption-lh); color: var(--text-tertiary);"
    >
      &copy; {new Date().getFullYear()}
      {tl('footer.copyright')}
    </span>
  </footer>
</div>
