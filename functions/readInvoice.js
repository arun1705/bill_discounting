'use strict';

var bcSdk = require('../fabcar/query');


/**
 * A module that will generate Invoice and store into the blockchain!
 * @module readInvoiceDetails
 */
/** Add Data into blockchain.*/
exports.readInvoiceDetails = (to) => {
    return new Promise((resolve, reject) => {

        bcSdk.readInvoiceDetails(to)
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



/**
 * A module that will generate Invoice and store into the blockchain!
 * @module readAllInvoiceDetails
 */
/** Add Data into blockchain.*/
exports.readAllInvoiceDetails = (startInvoiceId,endInvoiceId) => {
    return new Promise((resolve, reject) => {

        bcSdk.readAllInvoiceDetails(startInvoiceId,endInvoiceId)
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