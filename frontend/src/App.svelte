<script lang="ts">
  import logo from './assets/images/logo-universal.png';
  import { Greet, GetVersionInfo } from '../wailsjs/go/app/App';
  import { getTheme, toggleTheme, initTheme } from './lib/stores/theme';
  import { t, getLocale, setLocale } from './lib/i18n/index';
  import { Button, Card, IconButton } from './lib/components';

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
        <svg width="16" height="16" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path
            stroke-linecap="round"
            stroke-linejoin="round"
            stroke-width="2"
            d="M21 12a9 9 0 01-9 9m9-9a9 9 0 00-9-9m9 9H3m9 9a9 9 0 01-9-9m9 9c1.657 0 3-4.03 3-9s-1.343-9-3-9m0 18c-1.657 0-3-4.03-3-9s1.343-9 3-9m-9 9a9 9 0 019-9"
          />
        </svg>
      </IconButton>
      <IconButton variant="subtle" size="sm" label={tl('theme.toggle')} onclick={handleToggleTheme}>
        {#if currentTheme === 'dark'}
          <svg width="16" height="16" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2"
              d="M12 3v1m0 16v1m9-9h-1M4 12H3m15.364 6.364l-.707-.707M6.343 6.343l-.707-.707m12.728 0l-.707.707M6.343 17.657l-.707.707M16 12a4 4 0 11-8 0 4 4 0 018 0z"
            />
          </svg>
        {:else}
          <svg width="16" height="16" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2"
              d="M20.354 15.354A9 9 0 018.646 3.646 9.003 9.003 0 0012 21a9.003 9.003 0 008.354-5.646z"
            />
          </svg>
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
        <svg width="16" height="16" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path
            stroke-linecap="round"
            stroke-linejoin="round"
            stroke-width="2"
            d="M4 6a2 2 0 012-2h2a2 2 0 012 2v2a2 2 0 01-2 2H6a2 2 0 01-2-2V6zM14 6a2 2 0 012-2h2a2 2 0 012 2v2a2 2 0 01-2 2h-2a2 2 0 01-2-2V6zM4 16a2 2 0 012-2h2a2 2 0 012 2v2a2 2 0 01-2 2H6a2 2 0 01-2-2v-2zM14 16a2 2 0 012-2h2a2 2 0 012 2v2a2 2 0 01-2 2h-2a2 2 0 01-2-2v-2z"
          />
        </svg>
        {tl('actions.dashboard')}
      </Button>
      <Button variant="secondary" size="md" onclick={greet}>
        <svg width="16" height="16" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path
            stroke-linecap="round"
            stroke-linejoin="round"
            stroke-width="2"
            d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"
          />
        </svg>
        {tl('actions.scan')}
      </Button>
      <Button variant="outline" size="md" onclick={greet}>
        <svg width="16" height="16" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path
            stroke-linecap="round"
            stroke-linejoin="round"
            stroke-width="2"
            d="M10.325 4.317c.426-1.756 2.924-1.756 3.35 0a1.724 1.724 0 002.573 1.066c1.543-.94 3.31.826 2.37 2.37a1.724 1.724 0 001.066 2.573c1.756.426 1.756 2.924 0 3.35a1.724 1.724 0 00-1.066 2.573c.94 1.543-.826 3.31-2.37 2.37a1.724 1.724 0 00-2.573 1.066c-.426 1.756-2.924 1.756-3.35 0a1.724 1.724 0 00-2.573-1.066c-1.543.94-3.31-.826-2.37-2.37a1.724 1.724 0 00-1.066-2.573c-1.756-.426-1.756-2.924 0-3.35a1.724 1.724 0 001.066-2.573c-.94-1.543.826-3.31 2.37-2.37.996.608 2.296.07 2.572-1.065z"
          />
          <path
            stroke-linecap="round"
            stroke-linejoin="round"
            stroke-width="2"
            d="M15 12a3 3 0 11-6 0 3 3 0 016 0z"
          />
        </svg>
        {tl('actions.settings')}
      </Button>
    </div>

    <!-- Greet Result -->
    {#if resultText}
      <div
        class="mb-8 flex items-center gap-3"
        style="padding: 12px 16px; border-radius: var(--radius-md); background-color: var(--status-info-bg); border: 1px solid var(--status-info); font-size: var(--text-body); line-height: var(--text-body-lh); color: var(--status-info-text); box-shadow: var(--shadow-2);"
      >
        <svg
          width="16"
          height="16"
          fill="none"
          stroke="currentColor"
          viewBox="0 0 24 24"
          style="flex-shrink: 0; color: var(--status-info);"
        >
          <path
            stroke-linecap="round"
            stroke-linejoin="round"
            stroke-width="2"
            d="M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"
          />
        </svg>
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
            <svg
              width="20"
              height="20"
              fill="none"
              stroke="currentColor"
              viewBox="0 0 24 24"
              style="color: var(--accent-default);"
            >
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="2"
                d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z"
              />
            </svg>
            <span
              style="font-size: var(--text-caption); line-height: var(--text-caption-lh); font-weight: var(--weight-semibold); color: var(--text-primary); text-align: center;"
              >{tl('features.monitoring')}</span
            >
          </div>
        </Card>
        <Card variant="default" padding>
          <div class="flex flex-col items-center gap-2">
            <svg
              width="20"
              height="20"
              fill="none"
              stroke="currentColor"
              viewBox="0 0 24 24"
              style="color: var(--status-critical);"
            >
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="2"
                d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z"
              />
            </svg>
            <span
              style="font-size: var(--text-caption); line-height: var(--text-caption-lh); font-weight: var(--weight-semibold); color: var(--text-primary); text-align: center;"
              >{tl('features.threats')}</span
            >
          </div>
        </Card>
        <Card variant="default" padding>
          <div class="flex flex-col items-center gap-2">
            <svg
              width="20"
              height="20"
              fill="none"
              stroke="currentColor"
              viewBox="0 0 24 24"
              style="color: var(--accent-default);"
            >
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="2"
                d="M9 17v-2m3 2v-4m3 4v-6m2 10H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"
              />
            </svg>
            <span
              style="font-size: var(--text-caption); line-height: var(--text-caption-lh); font-weight: var(--weight-semibold); color: var(--text-primary); text-align: center;"
              >{tl('features.analytics')}</span
            >
          </div>
        </Card>
        <Card variant="default" padding>
          <div class="flex flex-col items-center gap-2">
            <svg
              width="20"
              height="20"
              fill="none"
              stroke="currentColor"
              viewBox="0 0 24 24"
              style="color: var(--status-attention);"
            >
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="2"
                d="M15 17h5l-1.405-1.405A2.032 2.032 0 0118 14.158V11a6.002 6.002 0 00-4-5.659V5a2 2 0 10-4 0v.341C7.67 6.165 6 8.388 6 11v3.159c0 .538-.214 1.055-.595 1.436L4 17h5m6 0v1a3 3 0 11-6 0v-1m6 0H9"
              />
            </svg>
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
