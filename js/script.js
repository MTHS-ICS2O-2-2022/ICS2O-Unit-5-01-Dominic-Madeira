// Copyright (c) 2023 Dominic M. All rights reserved
//
// Created by: Dominic M.
// Created on: April 2023
// This file contains the JS functions for index.html

"use strict"

const randomNumber = Math.floor(Math.random() * 6) + 1

/**
 * This function uses a selection component from https://github.com/CreativeIT/getmdl-select
 */

function myButtonClicked() {
  // input
  const numberGuessed = parseInt(document.getElementById('guessed-number').value)

  // process
  if (numberGuessed == randomNumber) {

  // output
  document.getElementById('answer').innerHTML = numberGuessed
}
