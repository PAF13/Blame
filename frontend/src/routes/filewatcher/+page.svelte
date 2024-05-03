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
		<div>
		
		  
		</div>
		<iframe
		id="inlineFrameExample"
		title="Inline Frame Example"
		width="1000"
		height="1000"
		src="https://eview.eplan.com">
	  </iframe>
	</div>
