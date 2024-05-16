<svelte:head>
	<title>About</title>
	<meta name="description" content="About this app" />
</svelte:head>
<script lang="ts">
	import { Label, Select,MultiSelect, Input,ButtonGroup,Button,Dropdown,Search,Checkbox,P} from 'flowbite-svelte';

	import {ImportVerbindungsliste} from "$lib/wailsjs/go/main/App";


	function dialog3(){
    ImportVerbindungsliste();
	}







</script>


<h1>Verbindungsliste</h1>


<ButtonGroup>
<Button color="dark" on:click={dialog3}>Export</Button>
</ButtonGroup>