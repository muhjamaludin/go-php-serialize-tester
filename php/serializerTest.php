<?php

use PHPUnit\Framework\TestCase;

require_once "serializer.php";

class SerializerTest extends TestCase
{
    public function testSerializeNumber()
    {
        $serializer = new Serializer();

        // write serialize php to json file
        $num = 901;
        $result = $serializer->serializerNumber($num);

        $phpSerialize = array(
            "base" => $num,
            "code" => "PHPSerializeNumber",
            "result" => $result
        );

        $path = "../go/result/phpSerializeNum.json";
        // Convert JSON data from an array to a string
        $jsonString = json_encode($phpSerialize, JSON_PRETTY_PRINT);
        // Write in the file
        $fp = fopen($path, 'w');
        fwrite($fp, $jsonString);
        fclose($fp);

        // read goSerializeNumber json
        $json = file_get_contents("../go/result/goSerializeNum.json");
        $json_data = json_decode($json, true);

        // get base data
        $TestBase = $json_data['base'];
        $serializer = new Serializer();


        // compare unserialize php to serialize go
        $this->assertEquals($TestBase, $serializer->unSerializerNumber($json_data['result']));
    }

    public function testSerializeString()
    {
        // write serialize php to json file
        $serializer = new Serializer;

        $str = "test string php complex 12345678";
        $result = $serializer->serializerString($str);

        $phpSerialize = array(
            "base" => $str,
            "code" => "PHPSerializeNumber",
            "result" => $result
        );

        $path = "../go/result/phpSerializeString.json";
        // Convert JSON data from an array to a string
        $jsonString = json_encode($phpSerialize, JSON_PRETTY_PRINT);
        // Write in the file
        $fp = fopen($path, 'w');
        fwrite($fp, $jsonString);
        fclose($fp);

        // read goSerializeString json
        $json = file_get_contents("../go/result/goSerializeString.json");
        $json_data = json_decode($json, true);

        // get base data
        $TestBase = $json_data['base'];
        $serializer = new Serializer();

        // compare unserialize php to serialize go
        $this->assertEquals($TestBase, $serializer->unSerializerString($json_data['result']));
    }
}
