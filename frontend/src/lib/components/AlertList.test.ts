import { describe, it, expect } from 'vitest';
import { render, screen } from '@testing-library/svelte';
import AlertList from './AlertList.svelte';
import type { Alert } from './AlertList.svelte';

const mockAlerts: Alert[] = [
  {
    id: '1',
    type: 'info',
    title: 'Info Alert',
    message: 'This is an info alert',
    time: '2m ago',
  },
  {
    id: '2',
    type: 'warning',
    title: 'Warning Alert',
    message: 'This is a warning',
    time: '5m ago',
  },
  {
    id: '3',
    type: 'critical',
    title: 'Critical Alert',
    message: 'This is critical',
    time: '10m ago',
  },
];

describe('AlertList', () => {
  it('renders all alerts', () => {
    render(AlertList, { props: { alerts: mockAlerts } });
    expect(screen.getByText('Info Alert')).toBeInTheDocument();
    expect(screen.getByText('Warning Alert')).toBeInTheDocument();
    expect(screen.getByText('Critical Alert')).toBeInTheDocument();
  });

  it('renders alert messages', () => {
    render(AlertList, { props: { alerts: mockAlerts } });
    expect(screen.getByText('This is an info alert')).toBeInTheDocument();
    expect(screen.getByText('This is a warning')).toBeInTheDocument();
    expect(screen.getByText('This is critical')).toBeInTheDocument();
  });

  it('renders alert times', () => {
    render(AlertList, { props: { alerts: mockAlerts } });
    expect(screen.getByText('2m ago')).toBeInTheDocument();
    expect(screen.getByText('5m ago')).toBeInTheDocument();
    expect(screen.getByText('10m ago')).toBeInTheDocument();
  });

  it('renders empty list', () => {
    const { container } = render(AlertList, { props: { alerts: [] } });
    expect(container.querySelectorAll('[style*="border-radius"]').length).toBe(0);
  });

  it('renders single alert', () => {
    render(AlertList, { props: { alerts: [mockAlerts[0]] } });
    expect(screen.getByText('Info Alert')).toBeInTheDocument();
    expect(screen.queryByText('Warning Alert')).not.toBeInTheDocument();
  });
});
