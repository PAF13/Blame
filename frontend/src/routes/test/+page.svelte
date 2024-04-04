<svelte:head>
	<title>About</title>
	<meta name="description" content="About this app" />
</svelte:head>
<script>
	import { Dialog, VerbindungCorrection } from "$lib/wailsjs/go/main/App";
	import { Label, Listgroup, ListgroupItem, ButtonGroup, Button, Input} from 'flowbite-svelte';

	let num = 0;


	function sliceAppending(){
		VerbindungCorrection();
	}
  </script>
  
  <div class="pt-8">
	<Label for="input-addon" class="mb-2">Alte Stückliste</Label>
	<ButtonGroup class="w-full">
		<Button color="dark" on:click={sliceAppending}>importieren</Button>
		<Input id="input-addon" type="number" value={num} placeholder="Dateipfad" />
	</ButtonGroup>
	</div>