import en from '../../locales/en.json';
import id from '../../locales/id.json';

type Messages = typeof en;

const messages: Record<string, Messages> = { en, id };

export type Locale = 'en' | 'id';

let current: Locale = (localStorage.getItem('locale') as Locale) || 'en';

export function getLocale(): Locale {
  return current;
}

export function setLocale(locale: Locale) {
  current = locale;
  localStorage.setItem('locale', locale);
  window.dispatchEvent(new CustomEvent('localechange', { detail: locale }));
}

function resolve(obj: unknown, keys: string[]): unknown {
  let val = obj;
  for (const k of keys) {
    if (val && typeof val === 'object') {
      val = (val as Record<string, unknown>)[k];
    } else {
      return undefined;
    }
  }
  return val;
}

export function t(key: string): string {
  const val = resolve(messages[current] || messages.en, key.split('.'));
  return typeof val === 'string' ? val : key;
}
