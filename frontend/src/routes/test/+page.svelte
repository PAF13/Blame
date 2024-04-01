<svelte:head>
	<title>About</title>
	<meta name="description" content="About this app" />
</svelte:head>
<script>
	import { Dialog, ExcelChoice } from "$lib/wailsjs/go/main/App";
	import { Listgroup, ListgroupItem, ButtonGroup, Button, Input} from 'flowbite-svelte';
	let trynum ="";
	let list1 = "";
	let list1Last = "replace";

	function dialog1(){
		if(list1 != list1Last){
			list1 = "loading..."
			list1Last = "loading..."
			Dialog().then((result) => (list1 = result));
		}
	}
  </script>
  
  
<ButtonGroup class="w-full">
	<Button color="dark" on:click={dialog1}>importieren</Button>
	<ListgroupItem>No files</ListgroupItem>
</ButtonGroup>

  