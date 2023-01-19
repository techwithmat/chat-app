<script lang="ts">
  import ChatHeader from './ChatHeader.svelte'
  import ChatFooter from './ChatFooter.svelte'
  import ChatBody from './ChatBody.svelte'
</script>

<div class="h-full grid grid-rows-[60px,1fr,60px] gap-4">
  <ChatHeader />
  <ChatBody />
  <ChatFooter />
</div>
