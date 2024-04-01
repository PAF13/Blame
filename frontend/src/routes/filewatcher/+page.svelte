<svelte:head>
	<title>About</title>
	<meta name="description" content="About this app" />
</svelte:head>

<script>
	import { page } from '$app/stores';
	import { goto } from '$app/navigation';
	function Goto(){
		goto('/filewatcher/neuesprojekt')
	}
</script>

<div class="text-column">
	<h1>File Watcher</h1>
</div>
