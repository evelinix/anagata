<script lang="ts">
  import { onMount } from 'svelte';

  interface Props {
    data: Array<{ time: string; value: number }>;
  }

  let { data }: Props = $props();

  let canvas: HTMLCanvasElement | undefined = $state(undefined);
  let container: HTMLDivElement | undefined = $state(undefined);

  function drawChart() {
    if (!canvas || !container || data.length === 0) return;

    const ctx = canvas.getContext('2d');
    if (!ctx) return;

    const rect = container.getBoundingClientRect();
    const dpr = window.devicePixelRatio || 1;
    const width = rect.width;
    const height = 200;

    canvas.width = width * dpr;
    canvas.height = height * dpr;
    canvas.style.width = `${width}px`;
    canvas.style.height = `${height}px`;
    ctx.scale(dpr, dpr);

    const padding = { top: 16, right: 16, bottom: 32, left: 48 };
    const chartW = width - padding.left - padding.right;
    const chartH = height - padding.top - padding.bottom;

    const values = data.map((d) => d.value);
    const maxVal = Math.max(...values, 1);
    const minVal = 0;

    ctx.clearRect(0, 0, width, height);

    const gridColor =
      getComputedStyle(document.documentElement).getPropertyValue('--stroke-subtle').trim() ||
      '#e5e7eb';
    const textColor =
      getComputedStyle(document.documentElement).getPropertyValue('--text-tertiary').trim() ||
      '#9ca3af';

    ctx.strokeStyle = gridColor;
    ctx.lineWidth = 1;
    ctx.setLineDash([4, 4]);
    for (let i = 0; i <= 4; i++) {
      const y = padding.top + (chartH / 4) * i;
      ctx.beginPath();
      ctx.moveTo(padding.left, y);
      ctx.lineTo(width - padding.right, y);
      ctx.stroke();
    }
    ctx.setLineDash([]);

    ctx.fillStyle = textColor;
    ctx.font = '11px Oxanium, sans-serif';
    ctx.textAlign = 'right';
    ctx.textBaseline = 'middle';
    for (let i = 0; i <= 4; i++) {
      const y = padding.top + (chartH / 4) * i;
      const val = Math.round(maxVal - (maxVal - minVal) * (i / 4));
      ctx.fillText(val.toString(), padding.left - 8, y);
    }

    ctx.textAlign = 'center';
    ctx.textBaseline = 'top';
    const step = Math.max(1, Math.floor(data.length / 6));
    for (let i = 0; i < data.length; i += step) {
      const x = padding.left + (chartW / (data.length - 1)) * i;
      ctx.fillText(data[i].time, x, height - padding.bottom + 8);
    }

    const accentColor =
      getComputedStyle(document.documentElement).getPropertyValue('--accent-default').trim() ||
      '#2563eb';
    ctx.beginPath();
    ctx.moveTo(padding.left, padding.top + chartH);
    for (let i = 0; i < data.length; i++) {
      const x = padding.left + (chartW / (data.length - 1)) * i;
      const y = padding.top + chartH - ((data[i].value - minVal) / (maxVal - minVal)) * chartH;
      ctx.lineTo(x, y);
    }
    ctx.lineTo(padding.left + chartW, padding.top + chartH);
    ctx.closePath();
    const gradient = ctx.createLinearGradient(0, padding.top, 0, padding.top + chartH);
    gradient.addColorStop(0, accentColor + '40');
    gradient.addColorStop(1, accentColor + '05');
    ctx.fillStyle = gradient;
    ctx.fill();

    ctx.beginPath();
    for (let i = 0; i < data.length; i++) {
      const x = padding.left + (chartW / (data.length - 1)) * i;
      const y = padding.top + chartH - ((data[i].value - minVal) / (maxVal - minVal)) * chartH;
      if (i === 0) ctx.moveTo(x, y);
      else ctx.lineTo(x, y);
    }
    ctx.strokeStyle = accentColor;
    ctx.lineWidth = 2;
    ctx.lineJoin = 'round';
    ctx.stroke();

    const lastX = padding.left + chartW;
    const lastY =
      padding.top + chartH - ((data[data.length - 1].value - minVal) / (maxVal - minVal)) * chartH;
    ctx.beginPath();
    ctx.arc(lastX, lastY, 4, 0, Math.PI * 2);
    ctx.fillStyle = accentColor;
    ctx.fill();
  }

  onMount(() => {
    drawChart();
    const observer = new ResizeObserver(() => drawChart());
    if (container) observer.observe(container);
    return () => observer.disconnect();
  });

  $effect(() => {
    void data;
    drawChart();
  });
</script>

<div bind:this={container} style="width: 100%; height: 200px;">
  <canvas bind:this={canvas} style="display: block;"></canvas>
</div>
