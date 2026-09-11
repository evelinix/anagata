<script lang="ts">
  import Typography from './Typography.svelte';

  export interface Alert {
    id: string;
    type: 'info' | 'warning' | 'critical';
    title: string;
    message: string;
    time: string;
  }

  interface Props {
    alerts: Alert[];
  }

  let { alerts }: Props = $props();

  const typeConfig = $derived.by(() => ({
    info: { color: 'var(--status-info)', bg: 'var(--status-info-bg)', icon: 'ℹ' },
    warning: { color: 'var(--status-attention)', bg: 'var(--status-attention-bg)', icon: '⚠' },
    critical: { color: 'var(--status-critical)', bg: 'var(--status-critical-bg)', icon: '✕' },
  }));
</script>

<div style="display: flex; flex-direction: column; gap: 8px;">
  {#each alerts as alert (alert.id)}
    {@const config = typeConfig[alert.type]}
    <div
      style="display: flex; gap: 12px; padding: 12px; border-radius: var(--radius-md); background-color: var(--bg-layer-default); border: 1px solid var(--stroke-subtle);"
    >
      <div
        style="width: 28px; height: 28px; border-radius: var(--radius-full); background-color: {config.bg}; display: flex; align-items: center; justify-content: center; flex-shrink: 0; color: {config.color}; font-size: 12px; font-weight: 600;"
      >
        {config.icon}
      </div>
      <div style="flex: 1; min-width: 0;">
        <div
          style="display: flex; align-items: center; justify-content: space-between; gap: 8px; margin-bottom: 2px;"
        >
          <Typography variant="body-strong">{alert.title}</Typography>
          <Typography variant="caption" color="tertiary" style="flex-shrink: 0;"
            >{alert.time}</Typography
          >
        </div>
        <Typography variant="body-small" color="secondary">{alert.message}</Typography>
      </div>
    </div>
  {/each}
</div>
