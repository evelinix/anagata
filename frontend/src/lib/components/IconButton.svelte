<script lang="ts">
  import type { Snippet } from 'svelte';

  type Variant = 'subtle' | 'outline' | 'transparent' | 'primary';
  type Size = 'sm' | 'md' | 'lg';

  let {
    variant = 'subtle',
    size = 'md',
    label = '',
    disabled = false,
    onclick,
    children,
  }: {
    variant?: Variant;
    size?: Size;
    label?: string;
    disabled?: boolean;
    onclick?: (e: MouseEvent) => void;
    children?: Snippet;
  } = $props();
</script>

<button
  class="winui-icon-btn winui-icon-btn--{variant} winui-icon-btn--{size}"
  class:winui-icon-btn--labeled={!!label}
  {disabled}
  aria-label={label}
  onclick={(e: MouseEvent) => {
    if (!disabled) onclick?.(e);
  }}
>
  {#if children}
    <span class="winui-icon-btn-icon">{@render children()}</span>
  {/if}
  {#if label}
    <span class="winui-icon-btn-label">{label}</span>
  {/if}
</button>

<style>
  .winui-icon-btn {
    display: inline-flex;
    align-items: center;
    justify-content: center;
    gap: 8px;
    border: none;
    cursor: pointer;
    font-family: var(--font-family);
    font-weight: var(--weight-semibold);
    border-radius: var(--radius-md);
    transition:
      background-color var(--duration-short) var(--easing-default),
      box-shadow var(--duration-short) var(--easing-default);
    outline: none;
    white-space: nowrap;
    user-select: none;
    color: var(--text-primary);
  }

  .winui-icon-btn:focus-visible {
    outline: 2px solid var(--accent-default);
    outline-offset: 2px;
  }

  .winui-icon-btn:disabled {
    opacity: 0.38;
    cursor: not-allowed;
    pointer-events: none;
  }

  .winui-icon-btn-icon {
    display: inline-flex;
    align-items: center;
    justify-content: center;
    flex-shrink: 0;
  }

  .winui-icon-btn-label {
    font-size: var(--text-body);
    line-height: var(--text-body-lh);
  }

  /* ---- Sizes ---- */
  .winui-icon-btn--sm {
    height: 28px;
    padding: 5px;
  }

  .winui-icon-btn--sm.winui-icon-btn--labeled {
    padding: 5px 10px;
  }

  .winui-icon-btn--md {
    height: 32px;
    padding: 6px;
  }

  .winui-icon-btn--md.winui-icon-btn--labeled {
    padding: 6px 14px;
  }

  .winui-icon-btn--lg {
    height: 40px;
    padding: 9px;
  }

  .winui-icon-btn--lg.winui-icon-btn--labeled {
    padding: 9px 18px;
  }

  /* ---- Subtle (default WinUI style) ---- */
  .winui-icon-btn--subtle {
    background-color: transparent;
  }

  .winui-icon-btn--subtle:hover {
    background-color: var(--stroke-subtle);
  }

  .winui-icon-btn--subtle:active {
    background-color: var(--stroke-subtle);
  }

  /* ---- Outline ---- */
  .winui-icon-btn--outline {
    background-color: transparent;
    border: 1px solid var(--stroke-default);
  }

  .winui-icon-btn--outline:hover {
    background-color: var(--stroke-subtle);
    border-color: var(--stroke-strong);
  }

  .winui-icon-btn--outline:active {
    background-color: var(--stroke-subtle);
  }

  /* ---- Transparent ---- */
  .winui-icon-btn--transparent {
    background-color: transparent;
  }

  .winui-icon-btn--transparent:hover {
    background-color: var(--stroke-subtle);
  }

  .winui-icon-btn--transparent:active {
    background-color: var(--stroke-subtle);
  }

  /* ---- Primary ---- */
  .winui-icon-btn--primary {
    background-color: var(--accent-fill);
    color: var(--text-on-accent);
    box-shadow: var(--shadow-2);
  }

  .winui-icon-btn--primary:hover {
    background-color: var(--accent-default);
    box-shadow: var(--shadow-4);
  }

  .winui-icon-btn--primary:active {
    background-color: var(--accent-dark);
    box-shadow: var(--shadow-0);
  }
</style>
