<script lang="ts">
	import { username } from '../stores'
	let name: string

	const handleSubmit = () => {
		username.set(name)
	}
</script>

<header class="mb-20">
	<h1 class="text-center text-zinc-100 text-5xl">Welcome to ChatApp.</h1>
</header>

<form
	class="w-full flex flex-col gap-2"
	on:submit|preventDefault={handleSubmit}
>
	<label for="user_input" class="text-zinc-200 text-xl">Username</label>
	<input
		type="text"
		placeholder="Enter your username"
		id="user_input"
		class="border border-zinc-500 rounded-md px-4 py-2 bg-zinc-700 outline-none text-zinc-200"
		bind:value={name}
	/>
	<button
		class="border border-transparent rounded-md p-2 text-lg bg-blue-500 hover:bg-blue-600 font-bold"
		type="submit">Connect</button
	>
</form>
