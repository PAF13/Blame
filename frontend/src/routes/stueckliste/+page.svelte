<svelte:head>
	<title>About</title>
	<meta name="description" content="About this app" />
</svelte:head>
<script>
	import { Dialog, ExcelChoice } from "$lib/wailsjs/go/main/App";
	let list1 = "";
	let list1Last = "replace";
	let list2 = "";
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
</script>
<div class="text-column">
	<h1>Stückliste</h1>
	<h6>Current: {list1}</h6>
	<button on:click={dialog1}>Stückliste alt</button>
	<h6>Current: {list2}</h6>
	<button on:click={dialog2}>Stückliste neu</button>
	<h3> </h3>
	<div>
		<button on:click={reset}>reset</button>
		<button on:click={compare}>vergleich</button>
	</div>
</div>
