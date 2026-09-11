<script lang="ts">
  import logo from './assets/images/logo-universal.png';
  import { GetVersionInfo } from '../wailsjs/go/app/App';
  import { getTheme, toggleTheme, initTheme } from './lib/stores/theme';
  import { t, getLocale, setLocale } from './lib/i18n/index';
  import {
    Button,
    Card,
    IconButton,
    Typography,
    StatusCard,
    AlertList,
    TrafficChart,
  } from './lib/components';
  import {
    IconWorld,
    IconSun,
    IconMoon,
    IconLayoutDashboard,
    IconSearch,
    IconSettings,
    IconChartBar,
    IconAlertTriangle,
    IconFileText,
    IconBell,
  } from './lib/components/icons';
  import { createMockData } from './lib/stores/dashboard';

  let page = $state<'home' | 'dashboard'>('home');
  let locale = $state(getLocale());
  let localeVersion = $state(0);
  let currentTheme = $state(getTheme());
  let versionInfo = $state<Record<string, string>>({});
  let activeTab = $state<'dashboard' | 'scan' | 'alerts' | 'settings'>('dashboard');
  let dashboardData = $state(createMockData());

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

  function refreshData() {
    dashboardData = createMockData();
  }

  type Tab = 'dashboard' | 'scan' | 'alerts' | 'settings';

  const navItems: Array<{ id: Tab; icon: typeof IconLayoutDashboard; labelKey: string }> = [
    { id: 'dashboard', icon: IconLayoutDashboard, labelKey: 'nav.dashboard' },
    { id: 'scan', icon: IconSearch, labelKey: 'nav.scan' },
    { id: 'alerts', icon: IconBell, labelKey: 'nav.alerts' },
    { id: 'settings', icon: IconSettings, labelKey: 'nav.settings' },
  ];
</script>

{#if page === 'home'}
  <!-- LANDING PAGE -->
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
        <IconButton
          variant="subtle"
          size="sm"
          label={tl('theme.toggle')}
          onclick={handleToggleTheme}
        >
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
        <img
          src={logo}
          class="w-16 h-16 mb-4"
          alt=""
          style="filter: drop-shadow(var(--shadow-4));"
        />
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
            <Typography variant="caption" color="tertiary" uppercase
              >{tl('status.system')}</Typography
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
        <Button variant="primary" size="md" onclick={() => (page = 'dashboard')}>
          <IconLayoutDashboard size={16} stroke={2} />
          {tl('actions.dashboard')}
        </Button>
        <Button
          variant="secondary"
          size="md"
          onclick={() => {
            page = 'dashboard';
            activeTab = 'scan';
          }}
        >
          <IconSearch size={16} stroke={2} />
          {tl('actions.scan')}
        </Button>
        <Button
          variant="outline"
          size="md"
          onclick={() => {
            page = 'dashboard';
            activeTab = 'settings';
          }}
        >
          <IconSettings size={16} stroke={2} />
          {tl('actions.settings')}
        </Button>
      </div>

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
              <Typography variant="caption-strong" align="center"
                >{tl('features.threats')}</Typography
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
              <Typography variant="caption-strong" align="center"
                >{tl('features.alerts')}</Typography
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
      <Typography variant="caption" color="tertiary">
        v{versionInfo.version || '0.1.0'}
      </Typography>
      <Typography variant="caption" color="tertiary">
        &copy; {new Date().getFullYear()}
        {tl('footer.copyright')}
      </Typography>
    </footer>
  </div>
{:else}
  <!-- DASHBOARD PAGE -->
  <div
    class="min-h-screen flex"
    style="background-color: var(--bg-mica); color: var(--text-primary);"
  >
    <!-- Sidebar -->
    <aside
      class="flex flex-col"
      style="width: 220px; border-right: 1px solid var(--stroke-subtle); background-color: var(--bg-layer-alt);"
    >
      <div
        class="flex items-center gap-3 px-4 py-4"
        style="border-bottom: 1px solid var(--stroke-subtle);"
      >
        <img src={logo} class="w-6 h-6" alt="" />
        <Typography variant="body-strong">{tl('app.title')}</Typography>
      </div>

      <nav class="flex-1 px-2 py-3" style="display: flex; flex-direction: column; gap: 2px;">
        {#each navItems as item (item.id)}
          <button
            class="flex items-center gap-3 px-3 py-2"
            style="border-radius: var(--radius-md); cursor: pointer; border: none; text-align: left; transition: background-color 0.15s; background-color: {activeTab ===
            item.id
              ? 'var(--bg-layer-active)'
              : 'transparent'}; color: {activeTab === item.id
              ? 'var(--text-primary)'
              : 'var(--text-secondary)'};"
            onclick={() => (activeTab = item.id)}
          >
            <item.icon size={18} stroke={2} />
            <Typography variant="body">{tl(item.labelKey)}</Typography>
          </button>
        {/each}
      </nav>

      <div class="px-4 py-3" style="border-top: 1px solid var(--stroke-subtle);">
        <div class="flex items-center gap-2 mb-2">
          <IconButton variant="subtle" size="sm" label="Toggle language" onclick={toggleLocale}>
            <IconWorld size={16} stroke={2} />
          </IconButton>
          <IconButton
            variant="subtle"
            size="sm"
            label={tl('theme.toggle')}
            onclick={handleToggleTheme}
          >
            {#if currentTheme === 'dark'}
              <IconSun size={16} stroke={2} />
            {:else}
              <IconMoon size={16} stroke={2} />
            {/if}
          </IconButton>
        </div>
        <Typography variant="caption" color="tertiary">v{versionInfo.version || '0.1.0'}</Typography
        >
      </div>
    </aside>

    <!-- Main Content -->
    <div class="flex-1 flex flex-col min-h-screen">
      <!-- Top bar -->
      <header
        class="flex items-center justify-between px-5 py-3"
        style="border-bottom: 1px solid var(--stroke-subtle); background-color: var(--bg-layer-default);"
      >
        <Typography variant="subtitle" weight="semibold">{tl(`nav.${activeTab}`)}</Typography>
        <div class="flex items-center gap-2">
          <IconButton variant="subtle" size="sm" label="Toggle language" onclick={toggleLocale}>
            <IconWorld size={16} stroke={2} />
          </IconButton>
          <IconButton
            variant="subtle"
            size="sm"
            label={tl('theme.toggle')}
            onclick={handleToggleTheme}
          >
            {#if currentTheme === 'dark'}
              <IconSun size={16} stroke={2} />
            {:else}
              <IconMoon size={16} stroke={2} />
            {/if}
          </IconButton>
          <Button variant="outline" size="sm" onclick={refreshData}>Refresh</Button>
        </div>
      </header>

      <!-- Content -->
      <main class="flex-1 px-5 py-5" style="overflow-y: auto;">
        {#if activeTab === 'dashboard'}
          <!-- Status Cards -->
          <div class="grid grid-cols-2 md:grid-cols-4 gap-3 mb-6">
            {#each dashboardData.monitors as monitor (monitor.label)}
              <StatusCard
                label={monitor.label}
                value={monitor.value}
                unit={monitor.unit}
                status={monitor.status}
              />
            {/each}
          </div>

          <!-- Two-column layout -->
          <div class="grid grid-cols-1 lg:grid-cols-3 gap-4">
            <!-- Traffic Chart (2/3 width) -->
            <div class="lg:col-span-2">
              <Card variant="default" padding>
                <div class="mb-3">
                  <Typography variant="body-strong">Network Traffic</Typography>
                  <Typography variant="caption" color="secondary">Last 24 hours</Typography>
                </div>
                <TrafficChart data={dashboardData.traffic} />
              </Card>
            </div>

            <!-- Recent Alerts (1/3 width) -->
            <Card variant="default" padding>
              <div class="mb-3">
                <Typography variant="body-strong">Recent Alerts</Typography>
                <Typography variant="caption" color="secondary"
                  >{dashboardData.alerts.length} alerts</Typography
                >
              </div>
              <div style="max-height: 240px; overflow-y: auto;">
                <AlertList alerts={dashboardData.alerts} />
              </div>
            </Card>
          </div>
        {:else if activeTab === 'scan'}
          <Card variant="default" padding>
            <div class="flex flex-col items-center gap-4 py-8">
              <IconSearch size={48} stroke={1.5} style="color: var(--accent-default);" />
              <Typography variant="title" as="h2">{tl('nav.scan')}</Typography>
              <Typography variant="body" color="secondary" align="center" style="max-width: 24rem;">
                Run a full system scan to detect potential threats and vulnerabilities.
              </Typography>
              <Button variant="primary" size="md" onclick={refreshData}>Start Scan</Button>
            </div>
          </Card>
        {:else if activeTab === 'alerts'}
          <Card variant="default" padding>
            <div class="mb-4">
              <Typography variant="body-strong">All Alerts</Typography>
              <Typography variant="caption" color="secondary"
                >{dashboardData.alerts.length} alerts</Typography
              >
            </div>
            <AlertList alerts={dashboardData.alerts} />
          </Card>
        {:else}
          <Card variant="default" padding>
            <div class="flex flex-col items-center gap-4 py-8">
              <IconSettings size={48} stroke={1.5} style="color: var(--accent-default);" />
              <Typography variant="title" as="h2">{tl('nav.settings')}</Typography>
              <Typography variant="body" color="secondary" align="center" style="max-width: 24rem;">
                Configure application settings and preferences.
              </Typography>
            </div>
          </Card>
        {/if}
      </main>

      <!-- Footer -->
      <footer
        class="flex items-center justify-between px-5 py-2"
        style="border-top: 1px solid var(--stroke-subtle); background-color: var(--bg-layer-default);"
      >
        <Typography variant="caption" color="tertiary">
          &copy; {new Date().getFullYear()}
          {tl('footer.copyright')}
        </Typography>
        <Typography variant="caption" color="tertiary">
          {versionInfo.os || ''} • {versionInfo.arch || ''}
        </Typography>
      </footer>
    </div>
  </div>
{/if}
