<svelte:head>
	<title>About</title>
	<meta name="description" content="About this app" />
</svelte:head>
<script>
	import { Dialog, VerbindungCorrection, XMLTest } from "$lib/wailsjs/go/main/App";
	import { Label, Listgroup, ListgroupItem, ButtonGroup, Button, Input} from 'flowbite-svelte';
	import {Dropdown, DropdownItem, Checkbox, Search } from 'flowbite-svelte';
  import { ChevronDownOutline, UserRemoveSolid } from 'flowbite-svelte-icons';

	let num = 0;


	function sliceAppending(){
		XMLTest();
	}

  </script>
  

<Button color="dark" on:click={sliceAppending}>Tauschen</Button>
