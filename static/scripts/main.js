new Vue({
    el: "#invoice-creator",
    data: () => ({
        valid: true,
        selected: "Norway",
        global: {
            number: "1234567890",
            ocrNumber: "9999999999",
            net: "",
            vat: "",
            gross: "",
            currency: "",
            invoiceDate: "2018-04-23",
            dueDate: "2018-04-24",
            issuer: {
                "orgNumber": "9898989898"
            },
            attention: "",
            ourReference: "",
            yourReference: "",
            customText: ""
        },
        customer: {
            name: '',      
            address1: "",
            address2: "",
            address3: "",
            postalCode: "",
            postalPlace: "",
            country: "",
            landline: "",
            mobile: "",
            email: ''
        },
        invoice: {
            number: "",
            articleNumber: "",
            text1: "",
            text2: "",
            text3: "",
            unit: "",
            quantity: "",
            unitPrice: "",
            discountRate: "",
            netPrice: "",
            vatRate: "",
            vatAmount: ""
        },

        invoices: [],
        rules: {
            nameRules: [
                v => !!v || 'Name is required'
            ],
            addressRules: [
                v => !!v || 'Address is required'
            ],
            emailRules: [
                v => !!v || 'E-mail is required',
                v => /^\w+([.-]?\w+)*@\w+([.-]?\w+)*(\.\w{2,3})+$/.test(v) || 'E-mail must be valid'
            ],
            phoneRules: [
                v => !!v || 'Phone number is required',
                v => /^[0-9]+$/.test(v) || 'Phone number must be valid'
            ],
            postalCodeRules: [
                v => !!v || 'Postal Code is required',
                v => /^[0-9]+$/.test(v) || 'Post Code must be valid'
            ],
            postalPlaceRules: [
                v => !!v || 'Postal Place is required',
                v => /^[a-zA-Z]+$/.test(v) || 'Postal Place must be valid'
            ],
            numberRules: [
                v => !!v || "Numbers only",
                v => /^[0-9]+$/.test(v) || 'Please use numbers only'
            ]
        },
        items: [
            { text: 'Norway' }
            ,
            { text: 'Sweden' },
            { text: 'Denmark' }
        ]
    }),
    watch: {

    },
    methods: {
        selectCountry(label) {
            if(label == "Norway"){
                this.customer.country = "NO";
                this.customer.landline = "+47";
                this.global.currency = "NOK";
            } else if(label == "Sweden"){
                this.customer.country = "SWE";
                this.customer.landline = "+46";
                this.global.currency = "SEK";
            } else {
                this.customer.country = "DK";
                this.customer.landline = "+45";
                this.global.currency = "DKK";
            }
            
        },
        submit() {
            if (this.$refs.form.validate()) {
                // Form submission to server is not yet supported
                let counter = 1;
                let json = this.$data.global;
                // VueJS reactivity would update delivery when we add number to customer object
                json.delivery = JSON.parse(JSON.stringify( this.$data.customer )); 
                json.customer = this.$data.customer;
                json.customer.number = "1234567890"; 
                
                json.lines = this.$data.invoices;
                json.lines.map(item => item.number = counter++);
                console.log(json);
            }
        },
        clear() {
            this.$refs.form.reset()
        },
        addInvoice() {
            this.invoices.push(Vue.util.extend({}, this.invoice))
        },
        removeInvoice(index) {
            Vue.delete(this.invoices, index);
            this.changeTotalNet();
            this.changeTotalVAT();
            this.ChangeGross();
        },
        changeNetPrice(index) {
            let currentInvoice = this.invoices[index];
            currentInvoice.netPrice = currentInvoice.quantity*currentInvoice.unitPrice-(currentInvoice.quantity*currentInvoice.unitPrice*(currentInvoice.discountRate/100));
            this.changeTotalNet();
        },
        changeVATAmount(index) {
            let currentInvoice = this.invoices[index];
            currentInvoice.vatAmount = currentInvoice.netPrice*(currentInvoice.vatRate/100)
            this.changeTotalVAT();
        },

        changeTotalNet() {
            this.global.net = this.invoices.reduce((total, item) => {
                return total + item.netPrice;
            }, 0);
            this.ChangeGross();
        },
        changeTotalVAT() {
            this.global.vat = this.invoices.reduce((total, item) => {
                return total + item.vatAmount;
            }, 0);
            this.ChangeGross();
        },
        ChangeGross() {
            this.global.gross = Number(this.global.net)+Number(this.global.vat);
        }
    }
})