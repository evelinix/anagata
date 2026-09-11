import { describe, it, expect } from 'vitest';
import { createMockData } from './dashboard';

describe('createMockData', () => {
  it('returns dashboard data with monitors', () => {
    const data = createMockData();
    expect(data.monitors).toHaveLength(4);
    expect(data.monitors[0].label).toBe('CPU');
    expect(data.monitors[1].label).toBe('Memory');
    expect(data.monitors[2].label).toBe('Network');
    expect(data.monitors[3].label).toBe('Disk');
  });

  it('returns dashboard data with alerts', () => {
    const data = createMockData();
    expect(data.alerts.length).toBeGreaterThan(0);
    expect(data.alerts[0]).toHaveProperty('id');
    expect(data.alerts[0]).toHaveProperty('type');
    expect(data.alerts[0]).toHaveProperty('title');
    expect(data.alerts[0]).toHaveProperty('message');
    expect(data.alerts[0]).toHaveProperty('time');
  });

  it('returns dashboard data with traffic', () => {
    const data = createMockData();
    expect(data.traffic).toHaveLength(24);
    expect(data.traffic[0]).toHaveProperty('time');
    expect(data.traffic[0]).toHaveProperty('value');
  });

  it('monitor values are within valid range', () => {
    const data = createMockData();
    for (const monitor of data.monitors) {
      expect(monitor.value).toBeGreaterThanOrEqual(0);
      expect(monitor.unit).toBeTruthy();
      expect(['normal', 'attention', 'critical']).toContain(monitor.status);
    }
  });

  it('alert types are valid', () => {
    const data = createMockData();
    for (const alert of data.alerts) {
      expect(['info', 'warning', 'critical']).toContain(alert.type);
    }
  });

  it('traffic values are non-negative', () => {
    const data = createMockData();
    for (const point of data.traffic) {
      expect(point.value).toBeGreaterThanOrEqual(0);
    }
  });
});
