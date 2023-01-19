<script lang="ts">
  // here I will set the user to null
  // set user connected
  let connected = 0
</script>

<div class="flex items-center justify-between border-b border-zinc-700">
  <button
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
  <p class="font-bold">
    <span class="text-blue-400">{connected}</span>
    {connected > 1 ? 'users' : 'user'} online
  </p>
</div>
