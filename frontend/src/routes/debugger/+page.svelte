<svelte:head>
	<title>Debugger</title>
	<meta name="description" content="About this app" />
</svelte:head>

<script>
	import { Label, Input, Button, InputAddon, ButtonGroup, Checkbox } from 'flowbite-svelte';
	import { OpenFileDialog } from "$lib/wailsjs/go/main/App";
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
			OpenFileDialog().then((result) => (list1 = result));
		}
	}
	function dialog2(){
		if(list2 != list2Last){
			list2 = "loading..."
			list2Last = "loading..."
			OpenFileDialog().then((result) => (list2 = result));
		}
	}
	function reset(){
		list1 = "";
		list2 = "";
	}

	function beistellung(){
		if (list1 && list2 != "loading..") {

		}
	}

	function switchList(){
		list1Switch = list1;
		list2Switch = list2;
		list2 = list1Switch;
		list1 = list2Switch;
	}

</script>

<div class="text-column">
	<h1>Debugger</h1>
	<br>
	<ButtonGroup class="w-full">
		<Button color="dark">Print Lagerbestand Siteca</Button>
		<Button color="dark" >Print Lagerbestand Kroenert</Button>	
		<Button color="dark">importieren</Button>	
		<Button color="dark">importieren</Button>			
	</ButtonGroup>
	<ButtonGroup class="w-full">
		<Button color="dark">Print Lagerbestand Siteca</Button>
		<Button color="dark">Print Lagerbestand Kroenert</Button>	
		<Button color="dark">importieren</Button>	
		<Button color="dark">importieren</Button>			
	</ButtonGroup>
</div>

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
		<Button color="dark" >Differenz</Button>
		<Button color="dark" on:click={beistellung}>Beistellung</Button>
	</div>
	