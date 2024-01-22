<?php

use PHPUnit\Framework\TestCase;

require_once "serializer.php";

class SerializerTest extends TestCase
{
    public function testSerializeNumber()
    {
        $serializer = new Serializer();

        $TestNumber = "12";
        $serializeNum = $serializer->serializerNumber($TestNumber);

        $this->assertEquals("i:12;", $serializeNum);
    }

    public function testSerializeString()
    {
        $serializer = new Serializer;

        $TestString = "test string complex 12345678";
        $serializeString = $serializer->serializerString($TestString);

        $this->assertEquals("s:23:\"test string complex 12345678\";", $serializeString);
    }
}
