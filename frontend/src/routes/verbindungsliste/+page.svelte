<svelte:head>
	<title>About</title>
	<meta name="description" content="About this app" />
</svelte:head>
<script>
	import { Dialog, LoadCSV } from "$lib/wailsjs/go/main/App";
	import { Label, Input, Button, InputAddon, ButtonGroup, Checkbox } from 'flowbite-svelte';
	let list1 = "";
	let list1Last = "replace";



	function dialog1(){
		if(list1 != list1Last){
			list1 = "loading..."
			list1Last = "loading..."
			Dialog().then((result) => (list1 = result));
		}
	}

	function reset(){
		list1 = "";
	}
	function compare(){
		if (list1 != "loading..") {
			LoadCSV(list1);
		}
	}
</script>

<h1>Verbindungsliste</h1>

<div class="pt-8">
<Label for="input-addon" class="mb-2">Current: {list1}</Label>
<ButtonGroup class="w-full">
	<Button color="dark" on:click={dialog1}>importieren</Button>
	<Input id="input-addon" type="email" placeholder="elonmusk" />
	
</ButtonGroup>
</div>
<div class="pt-8">
<ButtonGroup class="w-full">
	<Button color="dark" on:click={reset}>Reset</Button>
	<Button color="dark" on:click={compare}>Compare</Button>
	
</ButtonGroup>
</div>

