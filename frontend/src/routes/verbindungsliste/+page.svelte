<svelte:head>
	<title>About</title>
	<meta name="description" content="About this app" />
</svelte:head>
<script>
	import { Dialog, LoadCSV } from "$lib/wailsjs/go/main/App";
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
<div class="text-column">
	<h1>Verbindungsliste</h1>
	<h6>Current: {list1}</h6>
	<button on:click={dialog1}>Stückliste alt</button>
	<h3> </h3>
	<div>
		<button on:click={reset}>reset</button>
		<button on:click={compare}>vergleich</button>
	</div>
</div>
