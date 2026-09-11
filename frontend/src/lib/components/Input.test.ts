import { describe, it, expect } from 'vitest';
import { render, screen } from '@testing-library/svelte';
import Input from './Input.svelte';

describe('Input', () => {
  it('renders with label', () => {
    render(Input, { props: { label: 'Username', id: 'username' } });
    expect(screen.getByLabelText('Username')).toBeInTheDocument();
  });

  it('renders placeholder', () => {
    render(Input, { props: { placeholder: 'Enter text...' } });
    expect(screen.getByPlaceholderText('Enter text...')).toBeInTheDocument();
  });

  it('defaults to type text', () => {
    render(Input);
    expect(screen.getByRole('textbox')).toHaveAttribute('type', 'text');
  });

  it('supports type text', () => {
    render(Input, { props: { type: 'text' } });
    expect(screen.getByRole('textbox')).toHaveAttribute('type', 'text');
  });

  it('supports type password', () => {
    render(Input, { props: { type: 'password' } });
    expect(screen.getByDisplayValue('')).toHaveAttribute('type', 'password');
  });

  it('supports type email', () => {
    render(Input, { props: { type: 'email' } });
    expect(screen.getByRole('textbox')).toHaveAttribute('type', 'email');
  });

  it('supports type number', () => {
    render(Input, { props: { type: 'number' } });
    expect(screen.getByRole('spinbutton')).toHaveAttribute('type', 'number');
  });

  it('supports type search', () => {
    render(Input, { props: { type: 'search' } });
    expect(screen.getByRole('searchbox')).toHaveAttribute('type', 'search');
  });

  it('supports type url', () => {
    render(Input, { props: { type: 'url' } });
    expect(screen.getByRole('textbox')).toHaveAttribute('type', 'url');
  });

  it.each(['sm', 'md', 'lg'] as const)('applies %s size class', (size) => {
    const { container } = render(Input, { props: { size } });
    const input = container.querySelector('input');
    expect(input?.className).toContain(`winui-input--${size}`);
  });

  it('shows error message', () => {
    render(Input, { props: { error: 'Field is required' } });
    expect(screen.getByText('Field is required')).toBeInTheDocument();
    const input = screen.getByRole('textbox');
    expect(input.className).toContain('winui-input--error');
  });

  it('applies disabled state', () => {
    render(Input, { props: { disabled: true } });
    expect(screen.getByRole('textbox')).toBeDisabled();
  });

  it('applies readonly state', () => {
    render(Input, { props: { readonly: true } });
    expect(screen.getByRole('textbox')).toHaveAttribute('readonly');
  });

  it('passes name and id attributes', () => {
    render(Input, { props: { name: 'email', id: 'email-field' } });
    const input = screen.getByRole('textbox');
    expect(input).toHaveAttribute('name', 'email');
    expect(input).toHaveAttribute('id', 'email-field');
  });
});
