<svelte:head>
	<title>About</title>
	<meta name="description" content="About this app" />
</svelte:head>
<script>
	import { Dialog, ExcelChoice } from "$lib/wailsjs/go/main/App";
	let list1 = "";
	let list2 = "";
	function dialog1(){
		//Dialog().then((result) => (list1 = result));
		Dialog()
	}
	function dialog2(){
		//Dialog().then((result) => (list2 = result));
		Dialog()
	}
	function reset(){
		list1 = "";
		list2 = "";
	}
	function compare(){
		ExcelChoice(list1,list2);
	}
</script>
<div class="text-column">
	<h1>Stückliste</h1>
	<h3></h3>
	<button on:click={dialog1}>Stückliste alt</button>
	<h3></h3>
	<button on:click={dialog2}>Stückliste neu</button>
	<h3> </h3>
	<div>
		<button on:click={reset}>reset</button>
		<button on:click={compare}>vergleich</button>
	</div>
</div>
