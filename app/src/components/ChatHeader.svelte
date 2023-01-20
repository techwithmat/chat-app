<script lang="ts">
  export let connection: WebSocket
  import { username, messages } from '../stores'

  const handleButton = () => {
    username.set('')
    messages.set([])
    connection.close()
  }
</script>

<div class="flex items-center justify-between border-b border-zinc-700">
  <button
    type="button"
    on:click|once={handleButton}
    class="flex items-center py-1 px-2 border border-zinc-500 rounded-md bg-zinc-700 cursor-pointer"
  >
    <svg
      width="24px"
      height="24px"
      stroke-width="1.98"
      viewBox="0 0 24 24"
      fill="none"
      xmlns="http://www.w3.org/2000/svg"
      color="#ccc"
      ><path
        d="M15 6l-6 6 6 6"
        stroke="#ccc"
        stroke-width="1.98"
        stroke-linecap="round"
        stroke-linejoin="round"
      /></svg
    >
    <span class="text-gray-300">Leave Chat</span>
  </button>
</div>
