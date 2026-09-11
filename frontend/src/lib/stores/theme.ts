import { writable } from 'svelte/store';

type Theme = 'light' | 'dark';

function createThemeStore() {
  const stored = localStorage.getItem('theme') as Theme | null;
  const initial: Theme =
    stored || (window.matchMedia('(prefers-color-scheme: dark)').matches ? 'dark' : 'light');

  const { subscribe, set, update } = writable<Theme>(initial);

  return {
    subscribe,
    toggle: () => {
      update((current) => {
        const next = current === 'light' ? 'dark' : 'light';
        localStorage.setItem('theme', next);
        document.documentElement.classList.toggle('dark', next === 'dark');
        return next;
      });
    },
    set: (theme: Theme) => {
      localStorage.setItem('theme', theme);
      document.documentElement.classList.toggle('dark', theme === 'dark');
      set(theme);
    },
    init: () => {
      const current = initial;
      document.documentElement.classList.toggle('dark', current === 'dark');
    },
  };
}

export const theme = createThemeStore();
