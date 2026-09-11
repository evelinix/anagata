import { describe, it, expect } from 'vitest';

describe('Basic math', () => {
  it('adds two numbers', () => {
    expect(1 + 1).toBe(2);
  });

  it('multiplies two numbers', () => {
    expect(2 * 3).toBe(6);
  });
});

describe('String operations', () => {
  it('concatenates strings', () => {
    expect('hello' + ' ' + 'world').toBe('hello world');
  });

  it('converts to uppercase', () => {
    expect('anagata'.toUpperCase()).toBe('ANAGATA');
  });
});
