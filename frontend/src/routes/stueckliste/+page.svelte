<svelte:head>
	<title>About</title>
	<meta name="description" content="About this app" />
</svelte:head>
<script>
	import { Dialog, ExcelChoice } from "$lib/wailsjs/go/main/App";
	import { Label, Input, Button, InputAddon, ButtonGroup, Checkbox } from 'flowbite-svelte';
	let list1 = "";
	let list1Switch = "";
	let list1Last = "replace";
	let list2 = "";
	let list2Switch = "";
	let list2Last = "replace";


	function dialog1(){
		if(list1 != list1Last){
			list1 = "loading..."
			list1Last = "loading..."
			Dialog().then((result) => (list1 = result));
		}
	}
	function dialog2(){
		if(list2 != list2Last){
			list2 = "loading..."
			list2Last = "loading..."
			Dialog().then((result) => (list2 = result));
		}
	}
	function reset(){
		list1 = "";
		list2 = "";
	}
	function compare(){
		if (list1 && list2 != "loading..") {
			ExcelChoice(list1,list2);
		}
	}

	function switchList(){
		list1Switch = list1;
		list2Switch = list2;
		list2 = list1Switch;
		list1 = list2Switch;
	}
</script>

<h1>Stückliste</h1>
<div class="pt-8">
<Label for="input-addon" class="mb-2">Alte Stückliste</Label>
<ButtonGroup class="w-full">
	<Button color="dark" on:click={dialog1}>importieren</Button>
	<Input id="input-addon" type="text" value={list1} placeholder="Dateipfad" />
</ButtonGroup>
</div>
<div class="pt-8">
<Label for="input-addon" class="mb-2">Neue Stückliste</Label>
<ButtonGroup class="w-full">
	<Button color="dark" on:click={dialog2}>importieren</Button>
	<Input id="input-addon" type="text" value={list2} placeholder="Dateipfad" />
	
</ButtonGroup>
</div>

<div class="pt-8">
	<Button color="dark" on:click={switchList}>Tauschen</Button>
	<Button color="dark" on:click={reset}>Reset</Button>
	<Button color="dark" on:click={compare}>Compare</Button>
</div>

