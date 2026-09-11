import type { Alert } from '../components/AlertList.svelte';

export interface MonitorStatus {
  label: string;
  value: number;
  unit: string;
  status: 'normal' | 'attention' | 'critical';
}

export interface TrafficPoint {
  time: string;
  value: number;
}

export interface DashboardData {
  monitors: MonitorStatus[];
  alerts: Alert[];
  traffic: TrafficPoint[];
}

export function createMockData(): DashboardData {
  const now = new Date();
  const traffic: TrafficPoint[] = [];
  for (let i = 23; i >= 0; i--) {
    const d = new Date(now.getTime() - i * 3600000);
    traffic.push({
      time: d.toLocaleTimeString([], { hour: '2-digit', minute: '2-digit' }),
      value: Math.floor(Math.random() * 1000),
    });
  }

  return {
    monitors: [
      { label: 'CPU', value: 34, unit: '%', status: 'normal' },
      { label: 'Memory', value: 62, unit: '%', status: 'normal' },
      { label: 'Network', value: 128, unit: 'Mbps', status: 'attention' },
      { label: 'Disk', value: 48, unit: '%', status: 'normal' },
    ],
    alerts: [
      {
        id: '1',
        type: 'warning',
        title: 'High CPU usage',
        message: 'CPU usage exceeded 80% on worker-2',
        time: '2m ago',
      },
      {
        id: '2',
        type: 'info',
        title: 'Scan completed',
        message: 'Network scan finished, no threats detected',
        time: '15m ago',
      },
      {
        id: '3',
        type: 'critical',
        title: 'Suspicious traffic',
        message: 'Unusual outbound traffic detected on port 443',
        time: '1h ago',
      },
      {
        id: '4',
        type: 'info',
        title: 'System updated',
        message: 'Threat database updated to v2024.12',
        time: '3h ago',
      },
    ],
    traffic,
  };
}
