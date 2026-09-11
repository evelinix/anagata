import { describe, it, expect } from 'vitest';
import { render } from '@testing-library/svelte';
import Card from './Card.svelte';

describe('Card', () => {
  it('renders a card element', () => {
    const { container } = render(Card);
    expect(container.querySelector('.ants-card')).toBeInTheDocument();
  });

  it.each(['default', 'filled', 'outlined', 'acrylic'] as const)(
    'applies %s variant class',
    (variant) => {
      const { container } = render(Card, { props: { variant } });
      const card = container.querySelector('.ants-card');
      expect(card?.className).toContain(`ants-card--${variant}`);
    },
  );

  it('has padding by default', () => {
    const { container } = render(Card);
    const body = container.querySelector('.ants-card-body');
    expect(body?.className).toContain('ants-card-body--padded');
  });

  it('removes padding when padding=false', () => {
    const { container } = render(Card, { props: { padding: false } });
    const body = container.querySelector('.ants-card-body');
    expect(body?.className).not.toContain('ants-card-body--padded');
  });
});
