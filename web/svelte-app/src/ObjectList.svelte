<script>
import { onMount } from "svelte";
import { writable } from "svelte/store"
let object = {
name: "Test",
      x: 0,
      y: 0,
      z: 0,
};
let player = {
name: "Web",
      x: 0,
      y: 0,
      z: 0,
};
let objects = writable([]);
export let socket;

onMount(() => {
    if (socket.readyState !== 1) {
        socket = new WebSocket("ws://localhost:3000/api/");
    } else {
        return;
    }
});

socket.onmessage = (event) => {
    let obj = JSON.parse(event.data);
    objects.set([]);
    try {
        obj.forEach(element => {
                if (element.name !== "") {
                objects.update(arr => [...arr, element])
                }
        });
    }
    catch(e) {
        return;
    }
}
function fetchObjects() {
    let jsonPlayer = JSON.stringify(player);
    socket.send(jsonPlayer);
}

function sendObject() {
    console.log(socket.readyState);
    let name = document.getElementById("name").value;
    let x = document.getElementById("x").value;
    let y = document.getElementById("y").value;
    let z = document.getElementById("z").value;
    if (name === "") {
        return
    }
    object.name = name;
    object.x = parseFloat(x);
    object.y = parseFloat(y);
    object.z = parseFloat(z);
    let jsonObj = JSON.stringify(object);
    socket.send(jsonObj); 
    fetchObjects();
}

function goForward() {
    if (player.x == 0) {
        fetchObjects();
        player.x += 20;
    } else {
        player.x += 20;
        fetchObjects();
    }
}
function goBack() {
    if (player.x - 20 < 0) {
        player.x = 1;
    } else {
        player.x -= 20;
    }
    fetchObjects();
}

</script>


<div>
<h2>Objects</h2>
<br>
<br>
<label for="">Name: </label>
<input name="name" type="text" id="name" class="in"/>

<label for="x">X: </label>
<input name="x" type="text" id="x" value="0" class="in"/>

<label for="y">Y: </label>
<input name="y" type="text" id="y" value="0" class="in"/>

<label for="z">Z: </label>
<input name="z" type="text" id="z" value="0" class="in"/>
<br>
<button on:click={sendObject} class="send" value="send">send</button>
<br>
<button on:click={goBack} class="arrow">⬅</button>
<button on:click={goForward} class="arrow">➡</button>
<div class="list">

{#each $objects as {name, x, y, z}}
<div class="obj">

<h2>{name}</h2>
<p>x : {x}</p>
<p>y : {y}</p>
<p>z : {z}</p>
</div>
<br>
<br>
{/each}
</div>
</div>

<style>

.in {
    border-radius: 10px;
}
.in:hover {
    border-color: #ff3e00;
}
.obj {
    text-align: center;
    padding: 10px;
    width: 5em;
    outline: 3px;
    border-radius: 10px;
    border-style:solid;
    border-color:black;
    margin: 10px;
}
.list {
    margin-left: 30%;
    width: 40%;
    padding: 10px;
    display: flex;
    flex-wrap: wrap;
    justify-content: center;
    right: 1em;
    align-content: center;
}
.arrow {
    text-align: center;
    background: transparent;
    border-radius: 10px;
    transition: ease-in-out 0.2s;
}
.arrow:hover {
    color: #ff3e00;
    border-color: #ff3e00;
    font-weight: bold;
}
.send {
    border-radius: 10px;
    transition: ease-in-out 0.2s;
}
.send:hover {
    border-color: #ff3e00;
    color : #ff3e00;
    font-weight: bold;
}

label {
    font-size: 20px;
    color: #ff3e00;
    font-weight: bold;
}
</style>
