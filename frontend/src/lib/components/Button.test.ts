import { describe, it, expect, vi } from 'vitest';
import { render, screen } from '@testing-library/svelte';
import Button from './Button.svelte';

describe('Button', () => {
  it('renders a button element', () => {
    render(Button);
    expect(screen.getByRole('button')).toBeInTheDocument();
  });

  it.each(['primary', 'secondary', 'danger', 'ghost', 'outline', 'subtle'] as const)(
    'applies %s variant class',
    (variant) => {
      render(Button, { props: { variant } });
      expect(screen.getByRole('button').className).toContain(`winui-btn--${variant}`);
    },
  );

  it.each(['sm', 'md', 'lg'] as const)('applies %s size class', (size) => {
    render(Button, { props: { size } });
    expect(screen.getByRole('button').className).toContain(`winui-btn--${size}`);
  });

  it('calls onclick when clicked', async () => {
    const onClick = vi.fn();
    render(Button, { props: { onclick: onClick } });
    await screen.getByRole('button').click();
    expect(onClick).toHaveBeenCalledOnce();
  });

  it('does not call onclick when disabled', async () => {
    const onClick = vi.fn();
    render(Button, { props: { disabled: true, onclick: onClick } });
    const btn = screen.getByRole('button');
    expect(btn).toBeDisabled();
    await btn.click();
    expect(onClick).not.toHaveBeenCalled();
  });

  it('renders disabled attribute', () => {
    render(Button, { props: { disabled: true } });
    expect(screen.getByRole('button')).toBeDisabled();
  });
});
