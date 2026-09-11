import { describe, it, expect, vi } from 'vitest';
import { render, screen } from '@testing-library/svelte';
import IconButton from './IconButton.svelte';

describe('IconButton', () => {
  it('renders with aria-label', () => {
    render(IconButton, { props: { label: 'Close' } });
    expect(screen.getByRole('button', { name: 'Close' })).toBeInTheDocument();
  });

  it.each(['subtle', 'outline', 'transparent', 'primary'] as const)(
    'applies %s variant class',
    (variant) => {
      render(IconButton, { props: { variant, label: 'Test' } });
      expect(screen.getByRole('button').className).toContain(`winui-icon-btn--${variant}`);
    },
  );

  it.each(['sm', 'md', 'lg'] as const)('applies %s size class', (size) => {
    render(IconButton, { props: { size, label: 'Test' } });
    expect(screen.getByRole('button').className).toContain(`winui-icon-btn--${size}`);
  });

  it('calls onclick when clicked', async () => {
    const onClick = vi.fn();
    render(IconButton, { props: { label: 'Click', onclick: onClick } });
    await screen.getByRole('button').click();
    expect(onClick).toHaveBeenCalledOnce();
  });

  it('does not call onclick when disabled', async () => {
    const onClick = vi.fn();
    render(IconButton, { props: { label: 'Click', disabled: true, onclick: onClick } });
    const btn = screen.getByRole('button');
    expect(btn).toBeDisabled();
    await btn.click();
    expect(onClick).not.toHaveBeenCalled();
  });
});
