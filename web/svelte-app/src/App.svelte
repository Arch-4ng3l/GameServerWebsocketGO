<script>

    import { Router, Route, Link} from "svelte-routing"
    import { onMount } from "svelte";
    import ObjectList from "./ObjectList.svelte";
    import ConnList from "./ConnList.svelte";
    let socket;
    onMount(() => {
        socket = new WebSocket("ws://localhost:3000/api/");
    });
</script>

<main>
    <Router>
        <div class="navbar">
            <nav>
                <Link to="/objects">
                <div class="link">
                    <h3>Object</h3>
                </div>
                </Link>
                <Link to="/conn">
                <div class="link">
                    <h3>Connections</h3>
                </div>
                </Link>
            </nav>
        </div>
        <h1>Server Dashboard</h1>
        <Route path="/objects" >
            <ObjectList {socket}/>
        </Route>
        <Route path="/conn" component={ConnList}/>
        </Router>
</main>

<style>
	main {
		text-align: center;
		padding: 1em;
		max-width: 240px;
		margin: 0 auto;
	}

	h1 {
		color: #ff3e00;
		text-transform: uppercase;
		font-size: 4em;
		font-weight: 100;
	}

	@media (min-width: 640px) {
		main {
			max-width: none;
		}
	}


    .navbar{
        position: fixed;
        display: flexbox;
        width: 140px;
        height: auto;
        padding: 0;
        margin: 0;
        z-index: 9;
        overflow: hidden;
        box-shadow: 0 8px 30px 0 rgba(0,0,0,0.3);
        background-color: #353746;
        text-decoration: none;
    }
    .link {
        border-radius: 3px;
        border-width: 1px;
        border-style: solid;
        color: black;
        font-size: 10px;
        text-decoration: none;
        transition: ease-in-out 0.4s;
    }
    .link:hover{
        text-decoration: none;
        text-decoration-color: #ff3e00;
        border-width: 3px;
		color: #ff3e00;
        font-size: 12px;
        font-weight: bold;
    }
</style>
