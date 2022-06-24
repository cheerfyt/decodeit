<script lang="ts">
  import "./main.css";
  import { toast, SvelteToast } from "@zerodevx/svelte-toast";
  import { fly } from "svelte/transition";
  import { decodeEthRawTransaction } from "./utils/axios";

  const error = (msg: string) => {
    toast.push(msg, {
      theme: {
        "--toastBackground": "red",
        "--toastColor": "white",
        "--toastBarBackground": "olive",
        "--toastBorderRadius": ".2rem",
        "--toastBarHeight": ".1rem",
        "--toastMsgPadding": ".2rem",
        "--toastMinHeight": "3rem",
        "--toastPadding": "4px",
      },
      duration: 2000,
    });
  };

  let rawTransacion: string;
  let result: string;

  const showDonaction = () => {};

  const onCleanBtnClick = () => {
    rawTransacion = "";
    result = "";
  };

  const onDecodeBtnClick = () => {
    if (!rawTransacion || !rawTransacion.trim()) {
      return;
    }
    rawTransacion = rawTransacion.trim();
    result = "";

    decodeEthRawTransaction(rawTransacion)
      .then((res) => {
        if (res.data.error) {
          error(res.data.error);
        } else {
          result = JSON.stringify(res.data.tx, null, 4);
        }
      })
      .catch((err) => {
        error(err.message || err.stack);
      });
  };
</script>

<SvelteToast />

<main class="w-full h-screen flex flex-col justify-start items-center">
  <div
    class="w-full h-[40px] flex items-center justify-end bg-slate-500 shadow-2xl p-4">
    <button class="btn btn-primary btn-sm">donation</button>
  </div>

  <div class="w-1/2">
    <div class="flex flex-col gap-4">
      <textarea
        rows="20"
        cols="20"
        class="w-full h-48 textarea textarea-secondary text-accent-focus resize-none mt-12 p-4 font-mono font-medium text-lg bg-gray-300"
        placeholder="ETH like serialized transaction, support EIP1559 tranaction"
        bind:value="{rawTransacion}"></textarea>

      <div class="flex justify-between">
        <button
          class="{`btn btn-sm btn-secondary w-1/12 glass capitalize`}"
          on:click="{onDecodeBtnClick}">decoed</button>

        <button
          class="btn btn-sm btn-secondary w-1/12 glass capitalize"
          on:click="{onCleanBtnClick}">clean</button>
      </div>
    </div>

    {#if result}
      <div
        in:fly="{{ y: 100, duration: 1000 }}"
        out:fly="{{ x: 0, duration: 1000 }}">
        <div in:fly="{{ y: 50, duration: 500 }}" class="divider text-xl">
          RESULT
        </div>
        <div
          class="h-96 bg-gray-600 text-xl p-6 rounded-lg overflow-scroll whitespace-pre-wrap text-lime-300">
          <p>{result}</p>
        </div>
      </div>
    {/if}
  </div>
</main>
