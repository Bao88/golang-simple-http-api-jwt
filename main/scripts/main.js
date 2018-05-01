
// window.onload = function() {
//     let form = document.getElementById("loginForm");
//     form.addEventListener("submit", function(event){
//         event.preventDefault();
        
//     });

// }

// <button type="button" v-on:click="removeInvoice(index)" class="btn btn-block btn-danger">Remove</button>

// <v-text-field label="Article Number" v-model="invoice.articleNumber"></v-text-field>

// <v-text-field label="Description" v-model="invoice.text1"></v-text-field>

// <v-text-field label="Unit" v-model="invoice.unit"></v-text-field>

// <v-text-field label="Quantity" v-model="invoice.quantity" @change="changeNetPrice(index)" :rules="rules.numberRules"></v-text-field>

// <v-text-field label="Unit Price" v-model="invoice.unitPrice" @change="changeNetPrice(index)" :rules="rules.numberRules"></v-text-field>

// <v-text-field label="Discount Rate(%)" v-model="invoice.discountRate" @change="changeNetPrice(index)" :rules="rules.numberRules"></v-text-field>

// <v-text-field label="Net Price" v-model="invoice.netPrice" disabled></v-text-field>

// <v-text-field label="VAT Rate(%)" v-model="invoice.vatRate" @change="changeVATAmount(index)" :rules="rules.numberRules"></v-text-field>

// <v-text-field label="VAT Amount" v-model="invoice.vatAmount" disabled></v-text-field>

// <v-text-field label="Sum" v-bind:value="invoice.netPrice+invoice.vatAmount" disabled></v-text-field>

// <v-text-field label="Total Net Price" v-model="global.net" disabled></v-text-field>

// <v-text-field label="Total VAT Amount" v-model="global.vat" disabled></v-text-field>

// <v-text-field label="Gross" v-model="global.gross" disabled></v-text-field>


$(document).ready(function(){
    // alert("Hello");
    let form = $("#loginForm");

    $("#loginForm").on("submit", function(event){   
        event.preventDefault();
        // alert("submitting");
        $.ajax({
            type: "POST",
            url: "/login",
            data: $("#loginForm").serialize(),
            success: function(data){
                if(data == "Fail") alert("Wrong username/password!");
                else {
                    console.log(data);
                    localStorage.setItem("clientToken", data);
                    window.location.replace("/createInvoice");
                    // window.location.href = 'newPage.html';
                }
                // console.log("here");
            }
        })
    });

    $("#invoiceForm").on("submit", function(event){   
        event.preventDefault();
        // console.log($("#invoiceForm").serialize());
        // alert("submitting");
        $.ajax({
            type: "POST",
            url: "/invoice",
            data: $("#invoiceForm").serialize(),
            success: function(data){
                console.log(data);
            }
        })
    });

    $("#add").on("click", function(event){
        alert("add");   
    });
});