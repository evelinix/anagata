import { describe, it, expect, vi, beforeEach, afterEach } from 'vitest';
import { render } from '@testing-library/svelte';
import TrafficChart from './TrafficChart.svelte';

// Mock ResizeObserver + Canvas getContext
beforeEach(() => {
  vi.stubGlobal(
    'ResizeObserver',
    vi.fn(() => ({
      observe: vi.fn(),
      disconnect: vi.fn(),
      unobserve: vi.fn(),
    })),
  );
  vi.spyOn(HTMLCanvasElement.prototype, 'getContext').mockReturnValue({
    scale: vi.fn(),
    clearRect: vi.fn(),
    beginPath: vi.fn(),
    moveTo: vi.fn(),
    lineTo: vi.fn(),
    stroke: vi.fn(),
    fill: vi.fn(),
    arc: vi.fn(),
    closePath: vi.fn(),
    createLinearGradient: vi.fn(() => ({
      addColorStop: vi.fn(),
    })),
    fillText: vi.fn(),
    setLineDash: vi.fn(),
  } as unknown as CanvasRenderingContext2D);
});

afterEach(() => {
  vi.restoreAllMocks();
});

const mockData = [
  { time: '00:00', value: 100 },
  { time: '01:00', value: 200 },
  { time: '02:00', value: 150 },
  { time: '03:00', value: 300 },
];

describe('TrafficChart', () => {
  it('renders canvas element', () => {
    render(TrafficChart, { props: { data: mockData } });
    expect(document.querySelector('canvas')).toBeTruthy();
  });

  it('renders with empty data', () => {
    render(TrafficChart, { props: { data: [] } });
    expect(document.querySelector('canvas')).toBeTruthy();
  });

  it('renders container div', () => {
    const { container } = render(TrafficChart, { props: { data: mockData } });
    expect(container.querySelector('div')).toBeTruthy();
  });
});
