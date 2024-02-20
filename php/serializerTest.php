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

        // compare unserialize php to serialize go
        $this->assertEquals($TestBase, $serializer->unSerializerString($json_data['result']));
    }

    public function testSerializeArrayString()
    {
        // write serialize php to json file
        $serializer = new Serializer;

        $arr = array("name", "sheet 1", "Done", "", "1Ex");
        $result = $serializer->serializeArray($arr);

        $phpSerialize = array(
            "base" => $arr,
            "code" => "PHPSerializeArrayString",
            "result" => $result
        );

        $path = "../go/result/phpSerializeArrayString.json";
        // convert JSON data from an array to string
        $jsonString = json_encode($phpSerialize, JSON_PRETTY_PRINT);
        // write in the file
        $fp = fopen($path, 'w');
        fwrite($fp, $jsonString);
        fclose($fp);

        // read goSerializeArrayString json
        $json = file_get_contents("../go/result/goSerializeSliceString.json");
        $json_data = json_decode($json, true);

        // get base data
        $TestBase = $json_data['base'];
        // compare unserialize php to serialize go
        $this->assertEquals($TestBase, $serializer->unSerializeArray($json_data['result']));
    }

    public function testSerializeArrayInteger()
    {
        // write serialize php to json file
        $serializer = new Serializer;

        $arr = array(100, 200, 0, 1_000_000, 200_000);
        $result = $serializer->serializeArray($arr);

        $phpSerialize = array(
            "base" => $arr,
            "code" => "PHPSerializeArrayInteger",
            "result" => $result
        );

        $path = "../go/result/phpSerializeArrayInteger.json";
        // convert JSON data from an array to string
        $jsonString = json_encode($phpSerialize, JSON_PRETTY_PRINT);
        // write in the file
        $fp = fopen($path, 'w');
        fwrite($fp, $jsonString);
        fclose($fp);

        // read goserializeArrayInteger json
        $json = file_get_contents("../go/result/goSerializeSliceInt.json");
        $json_data = json_decode($json, true);

        // get base data
        $TestBase = $json_data['base'];
        // compare unserialize php to serialize go
        $this->assertEquals($TestBase, $serializer->unSerializeArray($json_data['result']));
    }

    public function testSerializeArrayStringInteger()
    {
        // write serialize php to json file
        $serializer = new Serializer;

        $arr = array("name", 100_000, "", "1Ex", 0, 1);
        $result = $serializer->serializeArray($arr);

        $phpSerialize = array(
            "base" => $arr,
            "code" => "PHPSerializeArrayStringInteger",
            "result" => $result
        );

        $path = "../go/result/phpSerializeArrayStringInteger.json";
        // convert JSON data from an array to string
        $jsonString = json_encode($phpSerialize, JSON_PRETTY_PRINT);
        // write in the file
        $fp = fopen($path, 'w');
        fwrite($fp, $jsonString);
        fclose($fp);

        // read goserializeArrayStringInteger json
        $json = file_get_contents("../go/result/goSerializeSliceStringInteger.json");
        $json_data = json_decode($json, true);

        // get base data
        $TestBase = $json_data['base'];
        // compare unserialize php to serialize go
        $this->assertEquals($TestBase, $serializer->unSerializeArray($json_data['result']));
    }

    public function testSerializeArrayStringIntegerBoolean()
    {
        // write serialize php to json file
        $serializer = new Serializer;

        $arr = array("name", true, "sheet 1", "Done", "1Ex", false);
        $result = $serializer->serializeArray($arr);

        $phpSerialize = array(
            "base" => $arr,
            "code" => "PHPSerializeArrayStringIntegerBoolean",
            "result" => $result
        );

        $path = "../go/result/phpSerializeArrayStringIntegerBoolean.json";
        // convert JSON data from an array to string
        $jsonString = json_encode($phpSerialize, JSON_PRETTY_PRINT);
        // write in the file
        $fp = fopen($path, 'w');
        fwrite($fp, $jsonString);
        fclose($fp);

        // read goSerializeArrayStringIntegerBoolean json
        $json = file_get_contents("../go/result/goSerializeSliceStringIntegerBoolean.json");
        $json_data = json_decode($json, true);

        // get base data
        $TestBase = $json_data['base'];
        // compare unserialize php to serialize go
        $this->assertEquals($TestBase, $serializer->unSerializeArray($json_data['result']));
    }
}
