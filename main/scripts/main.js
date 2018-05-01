$(document).ready(function(){
    $("#loginForm").on("submit", function(event){   
        event.preventDefault();
        $.ajax({
            type: "POST",
            url: "/login",
            data: $("#loginForm").serialize(),
            success: function(data){
                if(data == "Fail") alert("Wrong username/password!");
                else {
                    console.log(data);
                    localStorage.setItem("clientToken", data);
                    localStorage.setItem("user", $("#name").val());
                    window.location.replace("/createInvoice");
                }
            }
        })
    });
});