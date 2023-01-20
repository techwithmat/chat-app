<script lang="ts">
  export let connection: WebSocket
  let inputMessage: string

  const handleSendMessage = () => {
    if (inputMessage.length >= 3) {
      connection.send(
        JSON.stringify({ action: 'message', message: inputMessage })
      )
      inputMessage = ''
    }

    return
  }
</script>

<form
  class="flex items-center justify-between gap-2 border-t border-zinc-700"
  on:submit|preventDefault={handleSendMessage}
>
  <input
    type="text"
    placeholder="Write a message"
    id="user_input"
    bind:value={inputMessage}
    class="w-full border border-zinc-500 rounded-md px-3 py-2 bg-zinc-700 outline-none text-zinc-200"
  />
  <button
    class="border border-transparent rounded-md p-2 bg-blue-500 hover:bg-blue-600 font-bold"
    type="submit">Send</button
  >
</form>
