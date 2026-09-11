<script lang="ts">
  import type { Snippet } from 'svelte';

  type Variant = 'primary' | 'secondary' | 'danger' | 'ghost' | 'outline' | 'subtle';
  type Size = 'sm' | 'md' | 'lg';

  let {
    variant = 'primary',
    size = 'md',
    disabled = false,
    onclick,
    children,
  }: {
    variant?: Variant;
    size?: Size;
    disabled?: boolean;
    onclick?: (e: MouseEvent) => void;
    children?: Snippet;
  } = $props();
</script>

<button
  class="winui-btn winui-btn--{variant} winui-btn--{size}"
  {disabled}
  onclick={(e: MouseEvent) => {
    if (!disabled) onclick?.(e);
  }}
>
  {#if children}
    {@render children()}
  {/if}
</button>

<style>
  .winui-btn {
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
      box-shadow var(--duration-short) var(--easing-default),
      border-color var(--duration-short) var(--easing-default);
    outline: none;
    white-space: nowrap;
    user-select: none;
  }

  .winui-btn:focus-visible {
    outline: 2px solid var(--accent-default);
    outline-offset: 2px;
  }

  .winui-btn:disabled {
    opacity: 0.38;
    cursor: not-allowed;
    pointer-events: none;
  }

  /* ---- Sizes ---- */
  .winui-btn--sm {
    padding: 4px 12px;
    font-size: var(--text-body-small);
    line-height: var(--text-body-small-lh);
    height: 28px;
  }

  .winui-btn--md {
    padding: 5px 20px;
    font-size: var(--text-body);
    line-height: var(--text-body-lh);
    height: 32px;
  }

  .winui-btn--lg {
    padding: 8px 24px;
    font-size: var(--text-body);
    line-height: var(--text-body-lh);
    height: 40px;
  }

  /* ---- Primary ---- */
  .winui-btn--primary {
    background-color: var(--accent-fill);
    color: var(--text-on-accent);
    box-shadow: var(--shadow-2);
  }

  .winui-btn--primary:hover {
    background-color: var(--accent-default);
    box-shadow: var(--shadow-4);
  }

  .winui-btn--primary:active {
    background-color: var(--accent-dark);
    box-shadow: var(--shadow-0);
  }

  /* ---- Secondary ---- */
  .winui-btn--secondary {
    background-color: var(--bg-layer-alt);
    color: var(--text-primary);
    border: 1px solid var(--stroke-default);
    box-shadow: var(--shadow-2);
  }

  .winui-btn--secondary:hover {
    background-color: var(--stroke-subtle);
    border-color: var(--stroke-default);
  }

  .winui-btn--secondary:active {
    background-color: var(--stroke-subtle);
    box-shadow: var(--shadow-0);
  }

  /* ---- Danger ---- */
  .winui-btn--danger {
    background-color: var(--status-critical);
    color: rgb(255, 255, 255);
    box-shadow: var(--shadow-2);
  }

  .winui-btn--danger:hover {
    background-color: rgb(196, 43, 43);
    box-shadow: var(--shadow-4);
  }

  .winui-btn--danger:active {
    background-color: rgb(155, 33, 33);
    box-shadow: var(--shadow-0);
  }

  /* ---- Ghost ---- */
  .winui-btn--ghost {
    background-color: transparent;
    color: var(--text-primary);
  }

  .winui-btn--ghost:hover {
    background-color: var(--stroke-subtle);
  }

  .winui-btn--ghost:active {
    background-color: var(--stroke-subtle);
  }

  /* ---- Outline ---- */
  .winui-btn--outline {
    background-color: transparent;
    color: var(--text-primary);
    border: 1px solid var(--stroke-default);
  }

  .winui-btn--outline:hover {
    background-color: var(--stroke-subtle);
    border-color: var(--stroke-strong);
  }

  .winui-btn--outline:active {
    background-color: var(--stroke-subtle);
  }

  /* ---- Subtle ---- */
  .winui-btn--subtle {
    background-color: transparent;
    color: var(--text-secondary);
  }

  .winui-btn--subtle:hover {
    background-color: var(--stroke-subtle);
    color: var(--text-primary);
  }

  .winui-btn--subtle:active {
    background-color: var(--stroke-subtle);
  }
</style>
