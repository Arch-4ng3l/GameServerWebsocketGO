<script>
    import { auth, setCookie } from "./store";

    function signIn() {
       const username = document.getElementById("un").value;
       const password = document.getElementById("pw").value; 
       const obj = {
            username: username, 
            password: password,
       };

       const jsonObj = JSON.stringify(obj);
       let url = "/api/login"
       fetch(url, {
            method: "POST",
            body: jsonObj,
       })
       .then(response => {
            if(response.status === 200) {
                return response.json();
            }
       })
       .then(data => {
            const jwttoken = data.token;
            setCookie("token", jwttoken);
            $auth = true;
       });
        
    }
</script>


<div>
    <h2>Sign In</h2>
    <br>
    <label for="username"><b>Username</b></label> 
    <input class="in" id="un" type="text" name="username"/>
    <br>
    <br>
    <label for="passwd"><b>Password</b></label>
    <input class="in" id="pw" type="text" name="passwd"/>
    <br>
    <br>

    <button on:click={signIn} class="btn">Sign In</button>
</div>


<style>

    .in {
        border-radius: 10px;
    }
    .in:hover {
        border-color: #ff3e00;
    }
    .btn {
        border-radius: 10px;     
    }
    .btn:hover {
        border-color: #ff3e00;
        color: #ff3e00;
        font-weight: bold;
    }
</style>
