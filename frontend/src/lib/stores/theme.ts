type Theme = 'light' | 'dark';

function getInitial(): Theme {
  const stored = localStorage.getItem('theme') as Theme | null;
  if (stored === 'dark' || stored === 'light') return stored;
  return window.matchMedia('(prefers-color-scheme: dark)').matches ? 'dark' : 'light';
}

export function getTheme(): Theme {
  return getInitial();
}

export function toggleTheme(): Theme {
  const current = getInitial();
  const next: Theme = current === 'light' ? 'dark' : 'light';
  localStorage.setItem('theme', next);
  document.documentElement.classList.toggle('dark', next === 'dark');
  window.dispatchEvent(new CustomEvent('themechange', { detail: next }));
  return next;
}

export function initTheme() {
  const current = getInitial();
  document.documentElement.classList.toggle('dark', current === 'dark');
}
