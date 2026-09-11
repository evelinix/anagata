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
  class="ants-icon-btn ants-icon-btn--{variant} ants-icon-btn--{size}"
  {disabled}
  aria-label={label}
  onclick={(e: MouseEvent) => {
    if (!disabled) onclick?.(e);
  }}
>
  {#if children}
    <span class="ants-icon-btn-icon">{@render children()}</span>
  {/if}
</button>

<style>
  .ants-icon-btn {
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

  .ants-icon-btn:focus-visible {
    outline: 2px solid var(--accent-default);
    outline-offset: 2px;
  }

  .ants-icon-btn:disabled {
    opacity: 0.38;
    cursor: not-allowed;
    pointer-events: none;
  }

  .ants-icon-btn-icon {
    display: inline-flex;
    align-items: center;
    justify-content: center;
    flex-shrink: 0;
  }

  /* ---- Sizes ---- */
  .ants-icon-btn--sm {
    height: 28px;
    padding: 5px;
  }

  .ants-icon-btn--md {
    height: 32px;
    padding: 6px;
  }

  .ants-icon-btn--lg {
    height: 40px;
    padding: 9px;
  }

  /* ---- Subtle (default WinUI style) ---- */
  .ants-icon-btn--subtle {
    background-color: transparent;
  }

  .ants-icon-btn--subtle:hover {
    background-color: var(--stroke-subtle);
  }

  .ants-icon-btn--subtle:active {
    background-color: var(--stroke-subtle);
  }

  /* ---- Outline ---- */
  .ants-icon-btn--outline {
    background-color: transparent;
    border: 1px solid var(--stroke-default);
  }

  .ants-icon-btn--outline:hover {
    background-color: var(--stroke-subtle);
    border-color: var(--stroke-strong);
  }

  .ants-icon-btn--outline:active {
    background-color: var(--stroke-subtle);
  }

  /* ---- Transparent ---- */
  .ants-icon-btn--transparent {
    background-color: transparent;
  }

  .ants-icon-btn--transparent:hover {
    background-color: var(--stroke-subtle);
  }

  .ants-icon-btn--transparent:active {
    background-color: var(--stroke-subtle);
  }

  /* ---- Primary ---- */
  .ants-icon-btn--primary {
    background-color: var(--accent-fill);
    color: var(--text-on-accent);
    box-shadow: var(--shadow-2);
  }

  .ants-icon-btn--primary:hover {
    background-color: var(--accent-default);
    box-shadow: var(--shadow-4);
  }

  .ants-icon-btn--primary:active {
    background-color: var(--accent-dark);
    box-shadow: var(--shadow-0);
  }
</style>
