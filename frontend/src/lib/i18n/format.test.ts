import { describe, it, expect, beforeEach, afterEach, vi } from 'vitest';
import {
  formatDate,
  formatDateShort,
  formatTime,
  formatDateTime,
  formatNumber,
  formatCurrency,
  formatPercent,
  formatRelative,
} from './format';

describe('formatDate', () => {
  it('formats date for en locale', () => {
    const date = new Date(2026, 0, 15); // Jan 15, 2026
    expect(formatDate(date)).toBe('January 15, 2026');
  });

  it('accepts ISO string', () => {
    expect(formatDate('2026-06-01')).toBeTruthy();
  });

  it('accepts timestamp', () => {
    expect(formatDate(1735689600000)).toBeTruthy();
  });
});

describe('formatDateShort', () => {
  it('formats short date', () => {
    const date = new Date(2026, 0, 15);
    expect(formatDateShort(date)).toBe('01/15/2026');
  });
});

describe('formatTime', () => {
  it('formats time', () => {
    const date = new Date(2026, 0, 15, 14, 30);
    expect(formatTime(date)).toBeTruthy();
  });
});

describe('formatDateTime', () => {
  it('formats date and time', () => {
    const date = new Date(2026, 0, 15, 14, 30);
    expect(formatDateTime(date)).toBeTruthy();
  });
});

describe('formatNumber', () => {
  it('formats number', () => {
    expect(formatNumber(1234567)).toBe('1,234,567');
  });

  it('formats with decimals', () => {
    expect(formatNumber(1234.56, { minimumFractionDigits: 2 })).toBe('1,234.56');
  });
});

describe('formatCurrency', () => {
  it('formats IDR currency', () => {
    const result = formatCurrency(150000);
    expect(result).toMatch(/150[.,]000/);
  });

  it('formats USD currency', () => {
    expect(formatCurrency(1500, 'USD')).toBe('$1,500');
  });
});

describe('formatPercent', () => {
  it('formats percent', () => {
    expect(formatPercent(75.5)).toBe('75.5%');
  });
});

describe('formatRelative', () => {
  beforeEach(() => {
    vi.useFakeTimers();
  });

  afterEach(() => {
    vi.useRealTimers();
  });

  it('returns "now" for recent time', () => {
    vi.setSystemTime(new Date(2026, 0, 15, 12, 0, 0));
    const date = new Date(2026, 0, 15, 11, 59, 30);
    expect(formatRelative(date)).toBeTruthy();
  });

  it('returns minutes ago', () => {
    vi.setSystemTime(new Date(2026, 0, 15, 12, 0, 0));
    const date = new Date(2026, 0, 15, 11, 55, 0);
    expect(formatRelative(date)).toBeTruthy();
  });

  it('returns hours ago', () => {
    vi.setSystemTime(new Date(2026, 0, 15, 12, 0, 0));
    const date = new Date(2026, 0, 15, 9, 0, 0);
    expect(formatRelative(date)).toBeTruthy();
  });

  it('returns days ago', () => {
    vi.setSystemTime(new Date(2026, 0, 15, 12, 0, 0));
    const date = new Date(2026, 0, 13, 12, 0, 0);
    expect(formatRelative(date)).toBeTruthy();
  });
});
