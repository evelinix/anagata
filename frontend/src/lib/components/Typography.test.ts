import { describe, it, expect } from 'vitest';
import { render, screen } from '@testing-library/svelte';
import Typography from './Typography.svelte';

describe('Typography', () => {
  it('renders as paragraph by default', () => {
    render(Typography, { props: { text: 'Hello' } });
    expect(screen.getByText('Hello').tagName).toBe('P');
  });

  it('renders as specified element', () => {
    render(Typography, { props: { as: 'h1', text: 'Title' } });
    expect(screen.getByText('Title').tagName).toBe('H1');
  });

  it('renders as span', () => {
    render(Typography, { props: { as: 'span', text: 'Text' } });
    expect(screen.getByText('Text').tagName).toBe('SPAN');
  });

  it.each([
    'display',
    'title-large',
    'title',
    'subtitle',
    'body-large',
    'body',
    'body-small',
    'caption',
  ] as const)('applies %s variant size', (variant) => {
    render(Typography, { props: { variant, text: 'Text' } });
    const el = screen.getByText('Text');
    expect(el.style.fontSize).toBeTruthy();
  });

  it('applies color', () => {
    render(Typography, { props: { color: 'secondary', text: 'Text' } });
    const el = screen.getByText('Text');
    expect(el.style.color).toBeTruthy();
  });

  it('applies text alignment', () => {
    render(Typography, { props: { align: 'center', text: 'Text' } });
    const el = screen.getByText('Text');
    expect(el.style.textAlign).toBe('center');
  });

  it('applies uppercase', () => {
    render(Typography, { props: { uppercase: true, text: 'Text' } });
    const el = screen.getByText('Text');
    expect(el.style.textTransform).toBe('uppercase');
  });

  it('applies bold weight', () => {
    render(Typography, { props: { weight: 'bold', text: 'Text' } });
    const el = screen.getByText('Text');
    expect(el.style.fontWeight).toBeTruthy();
  });

  it('infers semibold for strong variants', () => {
    render(Typography, { props: { variant: 'body-strong', text: 'Text' } });
    const el = screen.getByText('Text');
    expect(el.style.fontWeight).toBeTruthy();
  });
});
