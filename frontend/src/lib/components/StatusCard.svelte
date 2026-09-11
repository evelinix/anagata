<script lang="ts">
  import type { Snippet } from 'svelte';
  import Typography from './Typography.svelte';

  interface Props {
    label: string;
    value: number;
    unit?: string;
    status?: 'normal' | 'attention' | 'critical';
    icon?: Snippet;
    class?: string;
    style?: string;
  }

  let {
    label,
    value,
    unit = '',
    status = 'normal',
    icon,
    class: className = '',
    style = '',
  }: Props = $props();

  const statusColor = $derived.by(() => {
    switch (status) {
      case 'critical':
        return 'var(--status-critical)';
      case 'attention':
        return 'var(--status-attention)';
      default:
        return 'var(--status-success)';
    }
  });

  const statusBg = $derived.by(() => {
    switch (status) {
      case 'critical':
        return 'var(--status-critical-bg)';
      case 'attention':
        return 'var(--status-attention-bg)';
      default:
        return 'var(--status-success-bg)';
    }
  });
</script>

<div class="ants-status-card {className}" {style}>
  <div style="display: flex; align-items: center; justify-content: space-between;">
    <Typography variant="caption" color="tertiary" uppercase>{label}</Typography>
    {#if icon}
      <div style="color: {statusColor};">{@render icon()}</div>
    {/if}
  </div>

  <div style="display: flex; align-items: baseline; gap: 4px;">
    <Typography variant="title" weight="bold" style="color: {statusColor};">{value}</Typography>
    {#if unit}
      <Typography variant="caption" color="secondary">{unit}</Typography>
    {/if}
  </div>

  <div style="height: 4px; border-radius: var(--radius-full); background-color: {statusBg};">
    <div
      style="height: 100%; width: {Math.min(
        value,
        100,
      )}%; border-radius: var(--radius-full); background-color: {statusColor}; transition: width 0.3s ease;"
    ></div>
  </div>
</div>

<style>
  .ants-status-card {
    background-color: var(--bg-layer-default);
    border: 1px solid var(--stroke-subtle);
    border-radius: var(--radius-lg);
    padding: 16px;
    display: flex;
    flex-direction: column;
    gap: 12px;
  }
</style>
