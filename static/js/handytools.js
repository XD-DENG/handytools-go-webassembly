function funcCopied() {
    alert('Copied.')
}

function produceEncodeDecodeClosure(resultTypeFlag, action, type) {
    return function () {
        this.seen = false;
        this.result_type_flag = resultTypeFlag;
        this.result = "... ...";

        if (this.value_to_encode_or_decode == "") {
            this.result = "";
        } else {
            let result = wasmEncodeDecode(
                this.value_to_encode_or_decode,
                action, type
            );
            if (result.startsWith("ERROR: ")) {
                this.message_color = "color:red";
            } else {
                this.message_color = "color:green";
                this.seen = true
            }
            this.result = result;
        }
    }
}

var generatePassword = new Vue({
    el: '#random-password',
    data: {
        numberlength: 3,
        lowerletterlength: 5,
        upperletterlength: 5,
        symbollength: 2,
        mustelements: "",
        message: "",
        message_color: 'color:black',
        seen: false
    },
    methods: {
        showRandomPassword: function () {
            this.seen = false;
            this.message = "... ...";
            let result = wasmGeneratePassword(
                parseInt(this.numberlength),
                parseInt(this.lowerletterlength),
                parseInt(this.upperletterlength),
                parseInt(this.symbollength),
                this.mustelements
            );
            if (result.startsWith("ERROR: ")) {
                this.message_color = "color:red";
            } else {
                this.message_color = "color:green";
                this.seen = true
            }
            this.message = result;
        },

        clear: function () {
            this.message = "";
            this.mustelements = "";
            this.seen = false
        },

        copied: funcCopied
    }
})

var hashCalculation = new Vue({
    el: '#hash-calculation',
    data: {
        quotation_placeholder: "",
        plain_value_to_apply: "",
        algorithm_to_apply: "md5",
        hashed_value: "",
        algo_for_hashed_value: "",
        plain_value_for_hashed_value: "",
        message_color: 'color:black',
        seen: false
    },
    methods: {
        calculateHash: function () {
            this.seen = false;
            this.quotation_placeholder = '"';
            this.algo_for_hashed_value = this.algorithm_to_apply;
            this.plain_value_for_hashed_value = this.plain_value_to_apply;
            this.hashed_value = "... ...";

            let result = wasmHashCalculation(
                this.plain_value_to_apply,
                this.algo_for_hashed_value
            );

            if (result.startsWith("ERROR: ")) {
                this.message_color = "color:red";
            } else {
                this.message_color = "color:green";
                this.seen = true;
            }
            this.hashed_value = result;

        },
        clear: function () {
            this.quotation_placeholder = "";
            this.plain_value_to_apply = "";
            this.hashed_value = "";
            this.algo_for_hashed_value = "";
            this.plain_value_for_hashed_value = ""
            this.seen = false
        },
        copied: funcCopied
    }
})

var urlEncodeDecode = new Vue({
    el: '#url-encode-decode',
    data: {
        value_to_encode_or_decode: "",
        result: "",
        result_type_flag: "",
        message_color: 'color:black',
        seen: false
    },
    methods: {
        urlEncode: produceEncodeDecodeClosure("Encoding ", "encode", "url"),
        urlDecode: produceEncodeDecodeClosure("Decoding ", "decode", "url"),
        clear: function () {
            this.value_to_encode_or_decode = "";
            this.result = "";
            this.result_type_flag = "";
            this.seen = false
        },
        copied: funcCopied
    }
})

var base64EncodeDecode = new Vue({
    el: '#base64-encode-decode',
    data: {
        value_to_encode_or_decode: "",
        result: "",
        result_type_flag: "",
        message_color: 'color:black',
        seen: false
    },
    methods: {
        base64Encode: produceEncodeDecodeClosure("Encoding ", "encode", "base64"),
        base64Decode: produceEncodeDecodeClosure("Decoding ", "decode", "base64"),
        clear: function () {
            this.value_to_encode_or_decode = "";
            this.result = "";
            this.result_type_flag = "";
            this.seen = false
        },
        copied: funcCopied
    }
})

var unixtime = new Vue({
    el: '#unixtime',
    data: {
        currentunixtime: "",
        unixtime: "",
        unixtimeprocessed: "",
        result: "",
        result_color: "color:black",
        seen: false
    },
    methods: {
        convertToTime: function () {

            this.seen = false;

            if (!parseInt(this.unixtime) && parseInt(this.unixtime) != 0) {
                this.unixtimeprocessed = this.unixtime;
                this.result_color = "color:red";
                this.result = "Invalid input";
            } else {
                this.unixtimeprocessed = parseInt(this.unixtime);
                this.result = "... ...";

                let result = wasmUnixTimeConverter(this.unixtimeprocessed) + " (" + wasmHumanReadableTimediff(parseInt(new Date().getTime() / 1000) - this.unixtimeprocessed) + ")";
                this.result_color = "color:green";
                this.result = result;
                this.seen = true
            }
        },

        updateCurrentUnixTime: function () {
            this.currentunixtime = parseInt(new Date().getTime() / 1000)
        },

        clear: function () {
            this.unixtime = "";
            this.unixtimeprocessed = "";
            this.result_color = "color:black";
            this.result = "";
            this.seen = false
        },

        copied: funcCopied
    }
})

setInterval(unixtime.updateCurrentUnixTime, 1000)
