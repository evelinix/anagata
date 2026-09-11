import { describe, it, expect } from 'vitest';
import { render, screen } from '@testing-library/svelte';
import StatusCard from './StatusCard.svelte';

describe('StatusCard', () => {
  it('renders label and value', () => {
    render(StatusCard, { props: { label: 'CPU', value: 34, unit: '%' } });
    expect(screen.getByText('CPU')).toBeInTheDocument();
    expect(screen.getByText('34')).toBeInTheDocument();
    expect(screen.getByText('%')).toBeInTheDocument();
  });

  it('renders without unit', () => {
    render(StatusCard, { props: { label: 'Status', value: 1 } });
    expect(screen.getByText('1')).toBeInTheDocument();
    expect(screen.queryByText('%')).not.toBeInTheDocument();
  });

  it.each(['normal', 'attention', 'critical'] as const)('renders status color for %s', (status) => {
    const { container } = render(StatusCard, {
      props: { label: 'CPU', value: 50, status },
    });
    const card = container.querySelector('.ants-status-card');
    expect(card).toBeInTheDocument();
  });

  it('renders without icon when not provided', () => {
    render(StatusCard, { props: { label: 'CPU', value: 50 } });
    expect(screen.getByText('CPU')).toBeInTheDocument();
  });

  it('applies custom class', () => {
    const { container } = render(StatusCard, {
      props: { label: 'CPU', value: 50, class: 'custom-class' },
    });
    const card = container.querySelector('.ants-status-card');
    expect(card?.className).toContain('custom-class');
  });
});
