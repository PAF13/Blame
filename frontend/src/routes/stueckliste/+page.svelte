<svelte:head>
	<title>About</title>
	<meta name="description" content="About this app" />
</svelte:head>
<script>
	import { Dialog, ExcelChoice,  ImportStueckliste, StuecklisteSum } from "$lib/wailsjs/go/main/App";
	import { Label, Input, Button, InputAddon, ButtonGroup, Checkbox } from 'flowbite-svelte';
	let list1 = "";
	let list1Switch = "";
	let list1Last = "replace";
	let list2 = "";
	let list2Switch = "";
	let list2Last = "replace";
	let list22 = Boolean(true);

	let startup = Boolean(false);
	let defaultModal = Boolean(false);

	function dialog1(){
		if(list1 != "loading..." && list2 != "loading..."){
			list1 = "loading..."
			
			Dialog().then((result) => (list1 = result));
		}
	}
	function dialog2(){
		if(list2 != "loading..." && list1 != "loading..."){
			list2 = "loading..."

			Dialog().then((result) => (list2 = result));
		}
	}
	function compare(){
		if (list1  != "loading..." || list2 != "loading...") {
			ExcelChoice(list1,list2);
		}
	}
	function compare2(){
		if (list1  != "loading..." || list2 != "loading...") {
			StuecklisteSum(list1);
		}
	}
	function reset(){
		list1 = "";
		list2 = "";
	}



	function switchList(){
		list1Switch = list1;
		list2Switch = list2;
		list2 = list1Switch;
		list1 = list2Switch;
	}



	$:if (list1 != "") {
		list22 = false
	
}

</script>
<h1>Stückliste</h1>
<h2>list1:       {list1}</h2>

<h2>list2:       {list2}</h2>

<h2>list22:      {list22}</h2>
<div class="pt-8">
<ButtonGroup class="w-full">
	<Button color="dark" on:click={dialog1}>importieren</Button>
	<Input id="input-addon" type="text" value={list1} placeholder="Alte Stückliste" />
	<Button color="dark" on:click={dialog2} disabled={list22}>importieren</Button>
	<Input id="input-addon" type="text" value={list2} placeholder="Neue Stückliste" />
</ButtonGroup>
</div>
<div class="pt-8">
<ButtonGroup>
	<Button color="dark" disabled={defaultModal} on:click={reset}>Reset</Button>
	<Button color="dark" disabled={defaultModal} on:click={switchList}>Tauschen</Button>
</ButtonGroup>
<ButtonGroup>
	<Button color="dark" disabled={defaultModal} on:click={compare}>Differenz</Button>	
	<Button color="dark" disabled={defaultModal} on:click={compare2}>Stückliste Sum</Button>
	<Button color="dark" disabled={defaultModal} on:click={compare}>Clean</Button>
</ButtonGroup>
</div>
