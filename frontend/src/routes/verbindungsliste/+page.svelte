<svelte:head>
	<title>About</title>
	<meta name="description" content="About this app" />
</svelte:head>
<script>
	import { Dialog, Verbindungsliste } from "$lib/wailsjs/go/main/App";
	import { Label, Input, Button, InputAddon, ButtonGroup, Checkbox } from 'flowbite-svelte';

	let list2 = "";

	let list2Last = "replace";



	function dialog2(){
		if(list2 != list2Last){
			list2 = "loading..."
			list2Last = "loading..."
			Dialog().then((result) => (list2 = result));
		}
	}
	function reset(){
		list2 = "";
	}
	function verbindungsliste(){
		if (list2 != "loading..") {
		Verbindungsliste(list2);
		}
	}

</script>

<div class="text-column">
	<h1>Verbindungsliste</h1>
</div>
<div class="pt-8">
	<Label for="input-addon" class="mb-2">Neue Stückliste</Label>
	<ButtonGroup class="w-full">
		<Button color="dark" on:click={dialog2}>importieren</Button>
		<Input id="input-addon" type="text" value={list2} placeholder="Dateipfad" />
		
	</ButtonGroup>
	</div>
	
	<div class="pt-8">
		<Button color="dark" on:click={reset}>Reset</Button>
		<Button color="dark" on:click={verbindungsliste}>Compare</Button>
	</div>

