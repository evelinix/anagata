<script lang="ts">
  import logo from './assets/images/logo-universal.png';
  import { Greet } from '../wailsjs/go/app/App';

  let resultText = $state('Please enter your name below 👇');
  let name = $state('');

  function greet() {
    Greet(name).then((result: string) => {
      resultText = result;
    });
  }
</script>

<div class="h-screen text-center">
  <img
    src={logo}
    class="block w-1/2 h-1/2 mx-auto pt-[10%] bg-center bg-no-repeat bg-[length:100%_100%] bg-content-box"
    alt="logo"
  />
  <div class="h-5 leading-5 my-6 mx-auto">{resultText}</div>
  <div class="flex items-center justify-center">
    <input
      class="h-[30px] leading-[30px] rounded-[3px] border-none outline-none px-[10px] bg-[rgba(240,240,240,1)] antialiased hover:bg-white focus:bg-white"
      bind:value={name}
      autocomplete="off"
      name="input"
      type="text"
    />
    <button
      class="w-[60px] h-[30px] leading-[30px] rounded-[3px] border-none ml-5 px-2 cursor-pointer hover:bg-gradient-to-t hover:from-[#cfd9df] hover:to-[#e2ebf0] hover:text-[#333333]"
      onclick={greet}
    >
      Greet
    </button>
  </div>
</div>
