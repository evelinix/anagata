<script lang="ts">
  import logo from './assets/images/logo-universal.png';
  import { Greet } from '../wailsjs/go/app/App';
  import { theme } from './lib/stores/theme';
  import { t, getLocale, setLocale } from './lib/i18n/index';

  let resultText = $state('');
  let name = $state('');
  let locale = $state(getLocale());
  let localeVersion = $state(0);

  $effect(() => {
    theme.init();
    const handler = (e: Event) => {
      locale = (e as CustomEvent).detail;
      localeVersion++;
    };
    window.addEventListener('localechange', handler);
    return () => window.removeEventListener('localechange', handler);
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

  function tl(key: string): string {
    void localeVersion;
    return t(key);
  }
</script>

<div
  class="h-screen text-center transition-colors duration-300 bg-white dark:bg-[rgba(27,38,54,1)] text-gray-900 dark:text-white"
>
  <div class="absolute top-4 right-4 flex gap-2">
    <button
      class="p-2 rounded-lg bg-gray-200 dark:bg-gray-700 hover:bg-gray-300 dark:hover:bg-gray-600 transition-colors text-sm"
      onclick={toggleLocale}
      aria-label="Toggle language"
    >
      {locale === 'en' ? 'ID' : 'EN'}
    </button>
    <button
      class="p-2 rounded-lg bg-gray-200 dark:bg-gray-700 hover:bg-gray-300 dark:hover:bg-gray-600 transition-colors"
      onclick={() => theme.toggle()}
      aria-label={tl('theme.toggle')}
    >
      {#if $theme === 'dark'}
        <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path
            stroke-linecap="round"
            stroke-linejoin="round"
            stroke-width="2"
            d="M12 3v1m0 16v1m9-9h-1M4 12H3m15.364 6.364l-.707-.707M6.343 6.343l-.707-.707m12.728 0l-.707.707M6.343 17.657l-.707.707M16 12a4 4 0 11-8 0 4 4 0 018 0z"
          />
        </svg>
      {:else}
        <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path
            stroke-linecap="round"
            stroke-linejoin="round"
            stroke-width="2"
            d="M20.354 15.354A9 9 0 018.646 3.646 9.003 9.003 0 0012 21a9.003 9.003 0 008.354-5.646z"
          />
        </svg>
      {/if}
    </button>
  </div>

  <img
    src={logo}
    class="block w-1/2 h-1/2 mx-auto pt-[10%] bg-center bg-no-repeat bg-[length:100%_100%] bg-content-box"
    alt={tl('app.title')}
  />
  <div class="h-5 leading-5 my-6 mx-auto">
    {#if resultText}
      {resultText}
    {:else}
      {tl('greet.placeholder')}
    {/if}
  </div>
  <div class="flex items-center justify-center">
    <input
      class="h-[30px] leading-[30px] rounded-[3px] border border-gray-300 dark:border-gray-600 outline-none px-[10px] bg-gray-100 dark:bg-gray-800 antialiased hover:bg-white dark:hover:bg-gray-700 focus:bg-white dark:focus:bg-gray-700 transition-colors"
      bind:value={name}
      autocomplete="off"
      name="input"
      type="text"
      placeholder={tl('greet.placeholder')}
    />
    <button
      class="w-[80px] h-[30px] leading-[30px] rounded-[3px] border-none ml-5 px-2 cursor-pointer bg-gray-200 dark:bg-gray-700 hover:bg-gradient-to-t hover:from-[#cfd9df] hover:to-[#e2ebf0] dark:hover:from-gray-600 dark:hover:to-gray-500 hover:text-[#333333] dark:hover:text-white transition-colors"
      onclick={greet}
    >
      {tl('greet.button')}
    </button>
  </div>
</div>
