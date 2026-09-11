<script lang="ts">
  import type { Snippet } from 'svelte';

  type Variant =
    | 'display'
    | 'title-large'
    | 'title'
    | 'title-strong'
    | 'subtitle'
    | 'subtitle-strong'
    | 'body-large'
    | 'body'
    | 'body-strong'
    | 'body-small'
    | 'caption'
    | 'caption-strong';

  type Color =
    | 'primary'
    | 'secondary'
    | 'tertiary'
    | 'disabled'
    | 'accent'
    | 'success'
    | 'critical'
    | 'attention';

  type As = 'p' | 'span' | 'h1' | 'h2' | 'h3' | 'h4' | 'h5' | 'h6' | 'div';

  let {
    variant = 'body',
    color = 'primary',
    as = 'p',
    align,
    weight,
    uppercase = false,
    text = '',
    children,
  }: {
    variant?: Variant;
    color?: Color;
    as?: As;
    align?: 'left' | 'center' | 'right';
    weight?: 'regular' | 'semibold' | 'bold';
    uppercase?: boolean;
    text?: string;
    children?: Snippet;
  } = $props();

  const sizeMap: Record<Variant, string> = {
    display: 'var(--text-display)',
    'title-large': 'var(--text-title-large)',
    title: 'var(--text-title)',
    'title-strong': 'var(--text-title-strong)',
    subtitle: 'var(--text-subtitle)',
    'subtitle-strong': 'var(--text-subtitle-strong)',
    'body-large': 'var(--text-body-large)',
    body: 'var(--text-body)',
    'body-strong': 'var(--text-body-strong)',
    'body-small': 'var(--text-body-small)',
    caption: 'var(--text-caption)',
    'caption-strong': 'var(--text-caption-strong)',
  };

  const lhMap: Record<Variant, string> = {
    display: 'var(--text-display-lh)',
    'title-large': 'var(--text-title-large-lh)',
    title: 'var(--text-title-lh)',
    'title-strong': 'var(--text-title-strong-lh)',
    subtitle: 'var(--text-subtitle-lh)',
    'subtitle-strong': 'var(--text-subtitle-strong-lh)',
    'body-large': 'var(--text-body-large-lh)',
    body: 'var(--text-body-lh)',
    'body-strong': 'var(--text-body-strong-lh)',
    'body-small': 'var(--text-body-small-lh)',
    caption: 'var(--text-caption-lh)',
    'caption-strong': 'var(--text-caption-strong-lh)',
  };

  const colorMap: Record<Color, string> = {
    primary: 'var(--text-primary)',
    secondary: 'var(--text-secondary)',
    tertiary: 'var(--text-tertiary)',
    disabled: 'var(--text-disabled)',
    accent: 'var(--accent-default)',
    success: 'var(--status-success-text)',
    critical: 'var(--status-critical-text)',
    attention: 'var(--status-attention-text)',
  };

  const weightMap: Record<string, string> = {
    regular: 'var(--weight-regular)',
    semibold: 'var(--weight-semibold)',
    bold: 'var(--weight-bold)',
  };

  function resolveWeight(): string {
    if (weight) return weightMap[weight];
    if (variant.includes('strong')) return 'var(--weight-semibold)';
    return 'var(--weight-regular)';
  }
</script>

<svelte:element
  this={as}
  style="
    font-size: {sizeMap[variant]};
    line-height: {lhMap[variant]};
    font-weight: {resolveWeight()};
    color: {colorMap[color]};
    {align ? `text-align: ${align};` : ''}
    {uppercase ? 'text-transform: uppercase; letter-spacing: 0.04em;' : ''}
  "
>
  {#if children}
    {@render children()}
  {:else if text}
    {text}
  {/if}
</svelte:element>
