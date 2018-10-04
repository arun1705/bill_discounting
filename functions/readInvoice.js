'use strict';

var bcSdk = require('../fabcar/query');


/**
 * A module that will generate Invoice and store into the blockchain!
 * @module readInvoiceDetails
 */
/** Add Data into blockchain.*/
exports.readInvoiceDetails = (to) => {
    return new Promise((resolve, reject) => {
        var resultArray = [];
        var invoiceids = [];
        var finalArray = [];
        var sorted_arr, uniqueItems;
        bcSdk.readInvoiceDetails(to)
            .then(function (result) {
                for (var i = 0; i < result.length; i++) {
                    resultArray.push(result[i].Value)
                }
                for (var i = 0; i < resultArray.length; i++) {
                    invoiceids.push(resultArray[i].invoiceID)
                }

                sorted_arr = invoiceids.slice().sort();
                uniqueItems = Array.from(new Set(sorted_arr))

                for (var i = 0; i < uniqueItems.length; i++) {
                    var index = resultArray.slice().reverse().findIndex(x => x['invoiceID'] === uniqueItems[i]);
                    var count = resultArray.length - 1
                    var finalIndex = index >= 0 ? count - index : index;
                    finalArray.push(resultArray[finalIndex])
                }
                resolve({
                    "message": finalArray
                });

            })
            .catch(err => {

                reject({
                    "status": 500,
                    "message": 'Something went wrong please try again later!!'
                });

            });
    });
}



/**
 * A module that will generate Invoice and store into the blockchain!
 * @module readAllInvoiceDetails
 */
/** Add Data into blockchain.*/
exports.readAllInvoiceDetails = (startInvoiceId, endInvoiceId) => {
    return new Promise((resolve, reject) => {

        bcSdk.readAllInvoiceDetails(startInvoiceId, endInvoiceId)
            .then(function (result) {
                console.log("result from read invoice--->>", result)
                resolve({
                    "status": 200,
                    "message": result
                });

            })
            .catch(err => {

                reject({
                    "status": 500,
                    "message": 'Something went wrong please try again later!!'
                });

            });
    });
}