<script lang="ts">
  export let user: string
  export let message: string
</script>

<div class="flex flex-col rounded-lg bg-zinc-800 border border-zinc-700 p-2">
  <p class="mb-2">
    <span class="font-bold">{user}</span> - Today
  </p>
  <p class="ml-2">{message}</p>
</div>
